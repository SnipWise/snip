package main

import (
	"context"
	"dmr-genkit-stream-completion/chatflow"
	"dmr-genkit-stream-completion/embeddings"
	"dmr-genkit-stream-completion/helpers"
	"dmr-genkit-stream-completion/rag"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/compat_oai/openai"
	"github.com/firebase/genkit/go/plugins/server"
	"github.com/openai/openai-go/option"
)

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

// SimilarityResult represents a single similarity search result
type SimilarityResult struct {
	ID         string  `json:"id"`
	Similarity float64 `json:"similarity"`
	Content    string  `json:"content"`
}

// SimilaritiesData holds the user message and found similarities
type SimilaritiesData struct {
	UserMessage string             `json:"user_message"`
	Count       int                `json:"count"`
	Results     []SimilarityResult `json:"results"`
}

// Global maps and mutexes
var (
	activeCompletions   = make(map[string]context.CancelFunc)
	completionsMutex    = sync.RWMutex{}
	pendingOperations   = make(map[string]*chatflow.OperationStatus)
	operationsMutex     = sync.RWMutex{}
	messages            []*ai.Message
	currentSimilarities SimilaritiesData
	similaritiesMutex   sync.RWMutex
)

func main() {
	ctx := context.Background()

	contextSizeLimit := helpers.StringToInt(helpers.GetEnvOrDefault("CONTEXT_SIZE_LIMIT", "3000"))

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

	embedder, store, err := embeddings.Generate(ctx, llmURL, embeddingsModel)
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
	streamingChatFlow := chatflow.DefineStreamingChatFlow(g, chatflow.StreamingChatFlowConfig{
		SnipModel:          snipModel,
		MemoryRetriever:    memoryRetriever,
		Messages:           &messages,
		ActiveCompletions:  &activeCompletions,
		CompletionsMutex:   &completionsMutex,
		PendingOperations:  &pendingOperations,
		OperationsMutex:    &operationsMutex,
		ContextSizeLimit:   contextSizeLimit,
		UpdateSimilarities: func(userMessage string, details []embeddings.SimilarityDetail) {
			similaritiesMutex.Lock()
			defer similaritiesMutex.Unlock()

			currentSimilarities.UserMessage = userMessage
			currentSimilarities.Count = len(details)
			currentSimilarities.Results = make([]SimilarityResult, len(details))
			for i, detail := range details {
				currentSimilarities.Results[i] = SimilarityResult{
					ID:         detail.ID,
					Similarity: detail.Similarity,
					Content:    detail.Content,
				}
			}
		},
	})

	mux := http.NewServeMux()
	if streamingChatFlow != nil {
		mux.HandleFunc("POST /completion", genkit.Handler(streamingChatFlow))
	}

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

	// Memory messages list endpoint
	mux.HandleFunc("GET /memory/messages/list", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Retrieving messages list")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status":   "ok",
			"messages": messages,
			"count":    len(messages),
		})
	})

	// Memory messages tokens endpoint
	mux.HandleFunc("GET /memory/messages/tokens", func(w http.ResponseWriter, r *http.Request) {
		totalTokens := helpers.CalculateTokenCount(messages)

		log.Printf("Calculated token count: %d", totalTokens)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status": "ok",
			"tokens": totalTokens,
			"count":  len(messages),
			"limit":  contextSizeLimit,
		})
	})

	// Similarities endpoint
	mux.HandleFunc("GET /similarities", func(w http.ResponseWriter, r *http.Request) {
		similaritiesMutex.RLock()
		similaritiesData := currentSimilarities
		similaritiesMutex.RUnlock()

		log.Printf("Retrieving similarities data: %d results found", similaritiesData.Count)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status": "ok",
			"data":   similaritiesData,
		})
	})

	// Models endpoint
	mux.HandleFunc("GET /models", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Retrieving models information")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status":           "ok",
			"chat_model":       snipModel,
			"embeddings_model": embeddingsModel,
		})
	})

	// Operation validate endpoint
	mux.HandleFunc("POST /operation/validate", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			OperationID string `json:"operation_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid request"})
			return
		}

		operationsMutex.Lock()
		operation, exists := pendingOperations[req.OperationID]
		if exists {
			operation.Status = "validated"
			operation.Continue <- true
		}
		operationsMutex.Unlock()

		if !exists {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "operation not found"})
			return
		}

		log.Printf("Operation %s validated", req.OperationID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status":       "ok",
			"operation_id": req.OperationID,
			"message":      "Operation validated and continued",
		})
	})

	// Operation cancel endpoint - marks as cancelled but continues
	mux.HandleFunc("POST /operation/cancel", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			OperationID string `json:"operation_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid request"})
			return
		}

		operationsMutex.Lock()
		operation, exists := pendingOperations[req.OperationID]
		if exists {
			operation.Status = "cancelled"
			operation.Continue <- true // Continue but with cancelled status
		}
		operationsMutex.Unlock()

		if !exists {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "operation not found"})
			return
		}

		log.Printf("Operation %s cancelled but continuing", req.OperationID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status":       "ok",
			"operation_id": req.OperationID,
			"message":      "Operation cancelled but stream continues",
		})
	})

	// Operation reset endpoint - marks as cancelled and stops
	mux.HandleFunc("POST /operation/reset", func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			OperationID string `json:"operation_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid request"})
			return
		}

		operationsMutex.Lock()
		operation, exists := pendingOperations[req.OperationID]
		if exists {
			operation.Status = "reset"
			operation.Continue <- false // Stop the stream
		}
		operationsMutex.Unlock()

		if !exists {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "operation not found"})
			return
		}

		log.Printf("Operation %s reset and stopped", req.OperationID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"status":       "ok",
			"operation_id": req.OperationID,
			"message":      "Operation reset and stopped",
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
