package chatflow

import (
	"context"
	"fmt"
	"log"
	"snip/embeddings"
	"strings"
	"sync"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// ChatRequest represents the input structure for the chat flow
type ChatRequest struct {
	Message string `json:"message"`
}

// ChatResponse represents the output structure for the chat flow
type ChatResponse struct {
	Response string `json:"response"`
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

// UpdateSimilaritiesFunc is a function type for updating similarities data
type UpdateSimilaritiesFunc func(userMessage string, details []embeddings.SimilarityDetail)

// OperationStatus represents the status of a pending operation
type OperationStatus struct {
	ID       string
	Status   string // "pending", "validated", "cancelled"
	Continue chan bool
}

// StreamingChatFlowConfig holds configuration for the streaming chat flow
type StreamingChatFlowConfig struct {
	SnipModel              string
	SystemInstruction      string
	ToolsModel             string
	ToolsSystemInstruction string
	MemoryRetriever        ai.Retriever
	Messages               *[]*ai.Message
	ActiveCompletions      *map[string]context.CancelFunc
	CompletionsMutex       *sync.RWMutex
	PendingOperations      *map[string]*OperationStatus
	OperationsMutex        *sync.RWMutex
	ContextSizeLimit       int
	Tools                  []ai.ToolRef
	UpdateSimilarities     UpdateSimilaritiesFunc
	//Genkit             *genkit.Genkit
}

// DefineStreamingChatFlow creates and returns a streaming chat flow
func DefineStreamingChatFlow(g *genkit.Genkit, config StreamingChatFlowConfig) *core.Flow[*ChatRequest, *ChatResponse, string] {
	return genkit.DefineStreamingFlow(
		g,
		"streaming-chat",
		func(ctx context.Context, input *ChatRequest, callback core.StreamCallback[string]) (*ChatResponse, error) {

			// [BEGIN] Tool calls detection
			// Create operation ID for this request
			operationID := fmt.Sprintf("op_%p", &ctx)

			// Create operation status and register it
			operation := &OperationStatus{
				ID:       operationID,
				Status:   "pending",
				Continue: make(chan bool, 1),
			}

			config.OperationsMutex.Lock()
			(*config.PendingOperations)[operationID] = operation
			config.OperationsMutex.Unlock()

			// Cleanup operation when done
			defer func() {
				config.OperationsMutex.Lock()
				delete(*config.PendingOperations, operationID)
				config.OperationsMutex.Unlock()
				close(operation.Continue)
			}()

			// [NOTE] deactivate and use ai.WithSystem and ai.WithPrompt

			// Initialize conversation history
			// history := []*ai.Message{}

			// history = append(history, ai.NewSystemTextMessage(config.ToolsSystemInstruction))
			// history = append(history, ai.NewUserTextMessage(input.Message))

			// Execute tool calls detection and execution
			toolResult, err := detectAndExecuteToolCalls(ctx, g, config, input.Message, operationID, operation, callback)
			if err != nil {
				return nil, err
			}

			if toolResult.TotalCalls > 0 {
				fmt.Println(strings.Repeat("+", 20))
				fmt.Println("🛠️ Total of tool calls made:", toolResult.TotalCalls)
				fmt.Println(strings.Repeat("+", 20))
				fmt.Println("🎉 Final response from the tools model:\n", toolResult.LastMessage)
				fmt.Println(strings.Repeat("+", 20))

				fmt.Println(strings.Repeat("*", 20))
				fmt.Println("🧮 Results:\n", toolResult.Results)
				fmt.Println(strings.Repeat("*", 20))

				*config.Messages = append(*config.Messages, ai.NewTextMessage("system", toolResult.Results))
			}
			// [END] Tool calls detection

			// [BEGIN] Similarity search
			err = performSimilaritySearch(ctx, input.Message, config)
			// [NOTE]: the function update directly the config.Messages
			// [TODO]: instead, return the list of similarities and append them here
			if err != nil {
				log.Fatal(err)
			}
			// [END] Similarity search

			// Update Conversational Memory
			// [NOTE] deactivate and use ai.WithSystem and ai.WithPrompt
			//*config.Messages = append(*config.Messages, ai.NewUserTextMessage(input.Message))

			// Debug: Print conversation state
			for i, msg := range *config.Messages {
				log.Printf("🟦  [%d] %s: %+v", i, msg.Role, msg.Content)
			}

			// Create a cancellable context for this completion
			completionCtx, cancel := context.WithCancel(ctx)

			// Generate a unique ID for this completion
			completionID := fmt.Sprintf("completion_%p", &completionCtx)

			// Store the cancel function
			config.CompletionsMutex.Lock()
			(*config.ActiveCompletions)[completionID] = cancel
			config.CompletionsMutex.Unlock()

			// Clean up when done
			defer func() {
				config.CompletionsMutex.Lock()
				delete(*config.ActiveCompletions, completionID)
				config.CompletionsMutex.Unlock()
			}()

			fmt.Println()
			fmt.Println(strings.Repeat("=", 80))
			fmt.Println()

			// [BEGIN] Stream Completion

			fullResponse, err := genkit.Generate(completionCtx, g,
				ai.WithModelName("openai/"+config.SnipModel),
				ai.WithSystem(config.SystemInstruction),

				ai.WithMessages(*config.Messages...),
				//ai.WithConfig(map[string]any{"temperature": 0.7}), // [NOTE]: set into the compose file

				ai.WithPrompt(input.Message),
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
			// [END] Stream Completion
			if err != nil {
				return nil, err
			}

			// Update Conversational Memory
			// Add assistant response to conversation history
			// [NOTE] deactivate and use ai.WithSystem and ai.WithPrompt
			//*config.Messages = append(*config.Messages, ai.NewTextMessage("assistant", fullResponse.Text()))

			*config.Messages = append(*config.Messages, ai.NewUserTextMessage(input.Message))
			*config.Messages = append(*config.Messages, ai.NewModelTextMessage(fullResponse.Text()))

			fmt.Println()
			fmt.Println(strings.Repeat("=", 80))
			fmt.Println()

			// Return the complete response for non-streaming clients
			return &ChatResponse{
				Response: fullResponse.Text(),
			}, nil
		},
	)
}
