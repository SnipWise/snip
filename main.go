package main

import (
	"context"
	"dmr-genkit-stream-completion/helpers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	fmt.Println("🌍 LLM URL:", llmURL)
	fmt.Println("🌍 SNIP Model:", snipModel)

	g := genkit.Init(ctx, genkit.WithPlugins(&openai.OpenAI{
		APIKey: "tada",
		Opts: []option.RequestOption{
			option.WithBaseURL(llmURL),
		},
	}))

	systemInstruction := helpers.GetEnvOrDefault("SYSTEM_INSTRUCTION", "You are a helpful AI assistant.")

	messages = append(messages, ai.NewSystemTextMessage(systemInstruction))

	// Definition of a streaming flow
	streamingChatFlow := genkit.DefineStreamingFlow(
		g,
		"streaming-chat",
		func(ctx context.Context, input *ChatRequest, callback core.StreamCallback[string]) (*ChatResponse, error) {

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
	log.Println("🤖 SNIP Server started and listening on: "+httpPort)
	log.Fatal(server.Start(ctx, "0.0.0.0:"+httpPort, mux))
}
