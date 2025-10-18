package chatflow

import (
	"context"
	"encoding/json"
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
	SnipModel          string
	ToolsModel         string
	MemoryRetriever    ai.Retriever
	Messages           *[]*ai.Message
	ActiveCompletions  *map[string]context.CancelFunc
	CompletionsMutex   *sync.RWMutex
	PendingOperations  *map[string]*OperationStatus
	OperationsMutex    *sync.RWMutex
	ContextSizeLimit   int
	Tools              []ai.ToolRef
	UpdateSimilarities UpdateSimilaritiesFunc
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

			// MEMORY:
			// STEP 1: Initialize conversation history, system message, and user message
			history := []*ai.Message{}

			// Create system message
			systemMsg := ``

			// STEP 2: Initialize loop control variables
			stopped := false           // Controls the conversation loop
			lastAssistantMessage := "" // Final AI message

			// STEP 3: Start the conversation loop
			// To avoid repeating the first user message in the history
			// we add it here before entering the loop and using prompt
			history = append(history, ai.NewUserTextMessage(input.Message))

			totalOfToolsCalls := 0

			// ai.NewUserTextMessage(input.Message)
			for !stopped {

				resp, err := genkit.Generate(ctx, g,
					ai.WithModelName("openai/"+config.ToolsModel),
					ai.WithSystem(systemMsg),
					// WithMessages sets the messages. These messages will be sandwiched between the system and user prompts.
					ai.WithMessages(history...),
					//ai.WithPrompt(userMsg),
					ai.WithTools(config.Tools...),
					ai.WithToolChoice(ai.ToolChoiceAuto),
					ai.WithReturnToolRequests(true),
				)

				if err != nil {
					fmt.Printf("🔴 [tools] Error: %v\n", err)
				}

				// We do not use parallel tool calls
				toolRequests := resp.ToolRequests()
				if len(toolRequests) == 0 {
					// No tool requests, we are done
					stopped = true // Exit the loop
					lastAssistantMessage = resp.Text()
					break // Exit the loop now
				}

				fmt.Println("✋ Number of tool requests", len(toolRequests))

				totalOfToolsCalls += len(toolRequests)

				// IMPORTANT: Add the assistant's message with tool requests to history
				// This ensures the model knows it already proposed these tools
				history = append(history, resp.Message)

				for _, req := range toolRequests {
					tool := genkit.LookupTool(g, req.Name)

					fmt.Println("🛠️ Tool request:", req.Name, req.Ref, req.Input)

					if tool == nil {
						log.Fatalf("tool %q not found", req.Name)
					}

					// Send pending status to client "kind":"tool_call",
					if callback != nil {

						inputJSON, err := json.Marshal(req.Input)
						if err != nil {
							return nil, fmt.Errorf("error marshaling tool input: %w", err)
						}
						inputJsonString := string(inputJSON)
						result := strings.ReplaceAll(inputJsonString, `"`, "")


						message := fmt.Sprintf("tool: %s %s", req.Name, result)
						//message := fmt.Sprintf("tool: %s", req.Name)
						pendingMsg := fmt.Sprintf(`{"kind":"tool_call", "message":"%s" , "status": "pending", "operation_id": "%s"}`, message, operationID)
						if err := callback(ctx, pendingMsg); err != nil {
							return nil, fmt.Errorf("error sending pending status: %w", err)
						}
					}

					log.Printf("⏸️  Operation %s waiting for confirmation...", operationID)

					// Wait for validation or cancellation
					select {
					case shouldContinue := <-operation.Continue:
						if !shouldContinue {
							log.Printf("❌ Operation %s cancelled by user", operationID)
							return nil, fmt.Errorf("operation cancelled by user")
						}
						log.Printf("✅ Operation %s validated, continuing...", operationID)
						output, err := tool.RunRaw(ctx, req.Input)
						if err != nil {
							log.Fatalf("tool %q execution failed: %v", tool.Name(), err)
						}
						fmt.Println("🤖 Result:", output)

						part := ai.NewToolResponsePart(&ai.ToolResponse{
							Name:   req.Name,
							Ref:    req.Ref,
							Output: output,
						})
						fmt.Println("✅", output)
						history = append(history, ai.NewMessage(ai.RoleTool, nil, part))

					case <-ctx.Done():
						log.Printf("⏱️  Operation %s context cancelled", operationID)
						return nil, ctx.Err()
					}

					fmt.Println(strings.Repeat("-", 20))
					fmt.Println("📜 History now has", len(history), "messages")
					fmt.Println(strings.Repeat("-", 20))

				}
			}

			if totalOfToolsCalls > 0 {
				fmt.Println(strings.Repeat("+", 20))
				fmt.Println("🛠️ Total of tool calls made:", totalOfToolsCalls)
				fmt.Println(strings.Repeat("+", 20))
				fmt.Println("🎉 Final response:\n", lastAssistantMessage)
				fmt.Println(strings.Repeat("+", 20))

				// [QUESTION]: assistant or system?
				// [TODO]: Check role
				// [TODO]: only if tool calls were made IMPORTANT:

				//*config.Messages = append(*config.Messages, ai.NewTextMessage("assistant", lastAssistantMessage))
				*config.Messages = append(*config.Messages, ai.NewTextMessage(
					"system",
					"TOOL CALLS RESULTS\n"+lastAssistantMessage+"\nEND OF TOOL CALLS RESULTS",
				))
			}

			// [END] Tool calls detection

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
