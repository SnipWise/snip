package chatflow

import (
	"context"
	"dmr-genkit-stream-completion/helpers"
	"dmr-genkit-stream-completion/rag"
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

// StreamingChatFlowConfig holds configuration for the streaming chat flow
type StreamingChatFlowConfig struct {
	SnipModel         string
	MemoryRetriever   ai.Retriever
	Messages          *[]*ai.Message
	ActiveCompletions *map[string]context.CancelFunc
	CompletionsMutex  *sync.RWMutex
}

// DefineStreamingChatFlow creates and returns a streaming chat flow
func DefineStreamingChatFlow(g *genkit.Genkit, config StreamingChatFlowConfig) *core.Flow[*ChatRequest, *ChatResponse, string] {
	return genkit.DefineStreamingFlow(
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
			retrieveResponse, err := config.MemoryRetriever.Retrieve(ctx, request)
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
				*config.Messages = append(*config.Messages, ai.NewSystemTextMessage("Relevant context:\n"+similarDocuments))
			}
			// -------------------------------------------------------------
			// [END] Similarity search
			// -------------------------------------------------------------

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
			if err != nil {
				return nil, err
			}

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
