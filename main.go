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
		SnipModel:         snipModel,
		MemoryRetriever:   memoryRetriever,
		Messages:          &messages,
		ActiveCompletions: &activeCompletions,
		CompletionsMutex:  &completionsMutex,
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
