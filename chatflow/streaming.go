package chatflow

import (
	"context"
	"dmr-genkit-stream-completion/embeddings"
	"fmt"
	"log"
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

// StreamingChatFlowConfig holds configuration for the streaming chat flow
type StreamingChatFlowConfig struct {
	SnipModel         string
	MemoryRetriever   ai.Retriever
	Messages          *[]*ai.Message
	ActiveCompletions *map[string]context.CancelFunc
	CompletionsMutex  *sync.RWMutex
	ContextSizeLimit  int
	UpdateSimilarities UpdateSimilaritiesFunc
}

// DefineStreamingChatFlow creates and returns a streaming chat flow
func DefineStreamingChatFlow(g *genkit.Genkit, config StreamingChatFlowConfig) *core.Flow[*ChatRequest, *ChatResponse, string] {
	return genkit.DefineStreamingFlow(
		g,
		"streaming-chat",
		func(ctx context.Context, input *ChatRequest, callback core.StreamCallback[string]) (*ChatResponse, error) {

			// [BEGIN] Similarity search
			// Retrieve relevant context from the vector store
			similarDocuments, details, err := embeddings.RetrieveSimilarDocuments(ctx, input.Message, config.MemoryRetriever)
			if err != nil {
				log.Fatal(err)
			}

			// Update global similarities data via callback
			if config.UpdateSimilarities != nil {
				config.UpdateSimilarities(input.Message, details)
			}

			if similarDocuments != "" {
				// Add Similarities to Messages
				*config.Messages = append(*config.Messages, ai.NewSystemTextMessage("Relevant context:\n"+similarDocuments))
			}
			// [END] Similarity search

			// Update Conversational Memory
			*config.Messages = append(*config.Messages, ai.NewUserTextMessage(input.Message))

			// Debug: Print conversation state
			for i, msg := range *config.Messages {
				log.Printf("  [%d] %s: %+v", i, msg.Role, msg.Content)
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
			
			// [BEGIN] Stream Completion
			fullResponse, err := genkit.Generate(completionCtx, g,
				ai.WithModelName("openai/"+config.SnipModel),

				ai.WithMessages(*config.Messages...),
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
			// [END] Stream Completion
			if err != nil {
				return nil, err
			}

			// Update Conversational Memory
			// Add assistant response to conversation history
			*config.Messages = append(*config.Messages, ai.NewTextMessage("assistant", fullResponse.Text()))

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
