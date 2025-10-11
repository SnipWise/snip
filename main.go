package main

import (
	"context"
	"dmr-genkit-stream-completion/helpers"
	"dmr-genkit-stream-completion/rag"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/compat_oai/openai"
	"github.com/firebase/genkit/go/plugins/server"
	"github.com/openai/openai-go/option"
)

// Structure for flow input
type ChatRequest struct {
	Message string `json:"message"`
}

// Structure for final flow output
type ChatResponse struct {
	Response string `json:"response"`
}

// ConversationMessage represents a single message in conversation history
type ConversationMessage struct {
	Role    string `json:"role"` // "user" or "assistant"
	Content string `json:"content"`
}

// Conversation holds the history for a conversation
type Conversation struct {
	ID       string                `json:"id"`
	Messages []ConversationMessage `json:"messages"`
}

// Global maps and mutexes
var (
	activeCompletions = make(map[string]context.CancelFunc)
	completionsMutex  = sync.RWMutex{}
	messages          []*ai.Message
)

func main() {
	ctx := context.Background()

	llmURL := helpers.GetEnvOrDefault("MODEL_RUNNER_BASE_URL", "http://localhost:12434/engines/llama.cpp/v1")
	snipModel := helpers.GetEnvOrDefault("SNIP_MODEL", "hf.co/menlo/jan-nano-gguf:q4_k_m")
	embeddingsModel := helpers.GetEnvOrDefault("EMBEDDING_MODEL", "ai/mxbai-embed-large:latest")

	fmt.Println("🌍 LLM URL:", llmURL)
	fmt.Println("🌍 SNIP Model:", snipModel)
	fmt.Println("🌍 Embeddings Model:", embeddingsModel)

	g := genkit.Init(ctx, genkit.WithPlugins(&openai.OpenAI{
		APIKey: "tada",
		Opts: []option.RequestOption{
			option.WithBaseURL(llmURL),
		},
	}))

	embedder, store, err := GenerateEmbeddings(ctx, llmURL, embeddingsModel)
	if err != nil {
		fmt.Println("😡 Error generating embeddings:", err)
		os.Exit(1)
	}
	// Create the memory vector retriever
	memoryRetriever := rag.DefineMemoryVectorRetriever(g, &store, embedder)
	fmt.Println("✅ Embeddings generated and vector store ready with", len(store.Records), "records")

	// [IMPORTANT] only for testing and checking
	fmt.Println("🚧 Example record:")
	fmt.Println(store.Records["9ce717f4-53ee-40f1-ac62-a3e0c55f67e4"].Prompt)

	systemInstruction := helpers.GetEnvOrDefault("SYSTEM_INSTRUCTION", "You are a helpful AI assistant.")

	messages = append(messages, ai.NewSystemTextMessage(systemInstruction))

	// Definition of a streaming flow
	streamingChatFlow := genkit.DefineStreamingFlow(
		g,
		"streaming-chat",
		func(ctx context.Context, input *ChatRequest, callback core.StreamCallback[string]) (*ChatResponse, error) {
			// -------------------------------------------------------------
			// [BEGIN] Similarity search
			// -------------------------------------------------------------
			// [IMPORTANT] Retrieve relevant context from the vector store
			// -------------------------------------------------------------
			// Use the custom retriever to find similar documents
			// -------------------------------------------------------------

			// Create a query document from the user question
			queryDoc := ai.DocumentFromText(input.Message, nil)

			similarityThreshold := helpers.StringToFloat(helpers.GetEnvOrDefault("SIMILARITY_THRESHOLD", "0.5"))
			similarityMaxResults := helpers.StringToInt(helpers.GetEnvOrDefault("SIMILARITY_MAX_RESULTS", "3"))
			// Create a retriever request with custom options
			request := &ai.RetrieverRequest{
				Query: queryDoc,
				Options: rag.MemoryVectorRetrieverOptions{
					Limit:      similarityThreshold,  // Lower similarity threshold to get more results
					MaxResults: similarityMaxResults, // Return top N results
				},
			}

			// Use the memory vector retriever to find similar documents
			retrieveResponse, err := memoryRetriever.Retrieve(ctx, request)
			if err != nil {
				log.Fatal(err)
			}

			similarDocuments := ""

			fmt.Printf("\nFound %d similar documents:\n", len(retrieveResponse.Documents))
			for i, doc := range retrieveResponse.Documents {
				similarity := doc.Metadata["cosine_similarity"]
				id := doc.Metadata["id"]
				fmt.Printf("%d. ID: %s, Similarity: %.4f\n", i+1, id, similarity)
				fmt.Printf("   Content: %s\n\n", doc.Content[0].Text)
				similarDocuments += doc.Content[0].Text
			}
			if similarDocuments != "" {
				messages = append(messages, ai.NewSystemTextMessage("Relevant context:\n"+similarDocuments))
			}
			// -------------------------------------------------------------
			// [END] Similarity search
			// -------------------------------------------------------------

			messages = append(messages, ai.NewUserTextMessage(input.Message))

			// Debug: Print conversation state
			for i, msg := range messages {
				log.Printf("  [%d] %s: %+v", i, msg.Role, msg.Content)
			}

			// Create a cancellable context for this completion
			completionCtx, cancel := context.WithCancel(ctx)

			// Generate a unique ID for this completion
			completionID := fmt.Sprintf("completion_%p", &completionCtx)

			// Store the cancel function
			completionsMutex.Lock()
			activeCompletions[completionID] = cancel
			completionsMutex.Unlock()

			// Clean up when done
			defer func() {
				completionsMutex.Lock()
				delete(activeCompletions, completionID)
				completionsMutex.Unlock()
			}()

			//

			fullResponse, err := genkit.Generate(completionCtx, g,
				ai.WithModelName("openai/"+snipModel),

				ai.WithMessages(messages...),
				//ai.WithConfig(map[string]any{"temperature": 0.7}),

				ai.WithStreaming(func(ctx context.Context, chunk *ai.ModelResponseChunk) error {
					// Do something with the chunk...
					fmt.Print(chunk.Text())

					if callback != nil {
						//fmt.Println("Sent chunk:", word)
						if err := callback(ctx, chunk.Text()); err != nil {
							return fmt.Errorf("error sending chunk: %w", err)
						}
					}

					return nil
				}),
			)
			if err != nil {
				return nil, err
			}

			// Add assistant response to conversation history
			messages = append(messages, ai.NewTextMessage("assistant", fullResponse.Text()))

			fmt.Println()
			fmt.Println(strings.Repeat("=", 80))
			fmt.Println()

			// Return the complete response for non-streaming clients
			return &ChatResponse{
				Response: fullResponse.Text(),
			}, nil
		},
	)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /completion", genkit.Handler(streamingChatFlow))

	// Stop completion endpoint
	mux.HandleFunc("POST /completion/stop", func(w http.ResponseWriter, r *http.Request) {
		completionsMutex.Lock()
		defer completionsMutex.Unlock()

		stoppedCount := 0
		for _, cancel := range activeCompletions {
			cancel()
			stoppedCount++
		}

		// Clear the map
		activeCompletions = make(map[string]context.CancelFunc)

		log.Printf("Stopped %d active completions", stoppedCount)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status":              "ok",
			"stopped_completions": stoppedCount,
			"message":             fmt.Sprintf("Stopped %d active completions", stoppedCount),
		})
	})

	// Memory reset endpoint
	mux.HandleFunc("POST /memory/reset", func(w http.ResponseWriter, r *http.Request) {
		if len(messages) > 0 {
			// Keep only the first system message
			systemMessage := messages[0]
			messages = []*ai.Message{systemMessage}
		}

		log.Println("Memory reset - kept only system message")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status":  "ok",
			"message": "Memory reset successfully",
		})
	})

	// Health endpoint
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		log.Println("Health check OK")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
	httpPort := helpers.GetEnvOrDefault("HTTP_PORT", "3500")
	log.Println("🤖 SNIP Server started and listening on: " + httpPort)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+httpPort, mux))
}

func GenerateEmbeddings(ctx context.Context, llmURL string, embeddingsModel string) (ai.Embedder, rag.MemoryVectorStore, error) {

	store := rag.MemoryVectorStore{
		Records: make(map[string]rag.VectorRecord),
	}

	oaiPlugin := &openai.OpenAI{
		APIKey: "tada",
		Opts: []option.RequestOption{
			option.WithBaseURL(llmURL),
		},
	}
	g := genkit.Init(ctx, genkit.WithPlugins(oaiPlugin))

	embedder := oaiPlugin.DefineEmbedder(embeddingsModel, nil)

	snippetsFolder := helpers.GetEnvOrDefault("SNIPPETS_FOLDER", "./snippets")

	_, err := helpers.ForEachFile(snippetsFolder, ".md", func(path string) error {
		jsonFilePath := strings.TrimSuffix(path, ".md") + ".json"
		fmt.Println("🚧 Path:", path)

		// Test if JSON file already exists
		if _, err := os.Stat(jsonFilePath); err == nil {
			fmt.Println("⚠️  JSON file already exists, skipping:", jsonFilePath)

			// [IMPORTANT] read the JSON file and load the vector record in the vector store
			data, err := os.ReadFile(jsonFilePath)
			if err != nil {
				fmt.Println("😡 Error reading JSON file:", err)
				return err
			}
			var record rag.VectorRecord
			if err := json.Unmarshal(data, &record); err != nil {
				fmt.Println("😡 Error unmarshaling JSON file:", err)
				return err
			}
			// Add the record to the vector store
			if _, err := store.Save(record); err != nil {
				fmt.Println("😡 Error saving vector record to store:", err)
				return err
			}

			//fmt.Println("✅ Loaded record from JSON file:", record.Prompt, record.Id)
			fmt.Println("✅ Loaded record from JSON file:", record.Id)

			return nil // Skip processing this file
		} else if !os.IsNotExist(err) {
			// Some other error
			fmt.Println("😡 Error checking JSON file:", err)
			return err
		} else {
			fmt.Println("🆕 JSON file does not exist, processing:", jsonFilePath)

			// Reading file content
			data, err := os.ReadFile(path)
			if err != nil {
				fmt.Println("😡 Error reading file:", err)
				return err
			}
			content := string(data)
			resp, err := genkit.Embed(ctx, g,
				ai.WithEmbedder(embedder),
				ai.WithTextDocs(content),
			)

			if err != nil {
				fmt.Println("😡 Error generating embedding:", err)
				return err
			}
			fmt.Println(resp.Embeddings)

			for i, emb := range resp.Embeddings {
				//fmt.Printf("Chunk %d (%s) embedding: %v\n", i, chunks[i], emb)

				// Store the embedding in the vector store
				record, errSave := store.Save(rag.VectorRecord{
					Prompt:    content,
					Embedding: emb.Embedding,
				})
				if errSave != nil {
					fmt.Println("😡 Error saving vector record:", errSave)
					return errSave
				}
				fmt.Println("-", i, "Saved record:", record.Prompt, record.Id)

				// Save the vector record to a JSON file
				vectorRecordJSON, err := json.MarshalIndent(record, "", "  ")
				if err != nil {
					fmt.Println("😡 Error marshaling vector record to JSON:", err)
					return err
				}

				// Write the JSON to a file
				err = os.WriteFile(jsonFilePath, vectorRecordJSON, 0644)
				if err != nil {
					fmt.Println("😡 Error writing vector record to file:", err)
					return err
				}
				return nil

			}

			return nil

		}

	})

	return embedder, store, err

}
