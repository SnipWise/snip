package chatflow

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/core"
	"github.com/firebase/genkit/go/genkit"
)

// toolCallsResult holds the result of tool calls detection and execution
type toolCallsResult struct {
	TotalCalls  int
	Results     string
	LastMessage string
}

// detectAndExecuteToolCalls handles the tool calls detection loop
func detectAndExecuteToolCalls(
	ctx context.Context,
	g *genkit.Genkit,
	config StreamingChatFlowConfig,
	history []*ai.Message,
	operationID string,
	operation *OperationStatus,
	callback core.StreamCallback[string],
) (*toolCallsResult, error) {
	stopped := false
	lastToolAssistantMessage := ""
	totalOfToolsCalls := 0
	toolCallsResults := ""

	for !stopped {
		resp, err := genkit.Generate(ctx, g,
			ai.WithModelName("openai/"+config.ToolsModel),
			ai.WithMessages(history...),
			ai.WithTools(config.Tools...),
			ai.WithToolChoice(ai.ToolChoiceAuto),
			ai.WithReturnToolRequests(true),
		)

		if err != nil {
			fmt.Printf("🔴 [tools] Error: %v\n", err)
		}

		toolRequests := resp.ToolRequests()
		if len(toolRequests) == 0 {
			stopped = true
			lastToolAssistantMessage = resp.Text()
			break
		}

		fmt.Println("✋ Number of tool requests", len(toolRequests))
		totalOfToolsCalls += len(toolRequests)
		history = append(history, resp.Message)

		for _, req := range toolRequests {
			tool := genkit.LookupTool(g, req.Name)

			fmt.Println("🛠️ Tool request:", req.Name, req.Ref, req.Input)

			if tool == nil {
				log.Fatalf("tool %q not found", req.Name)
			}

			if callback != nil {
				inputJSON, err := json.Marshal(req.Input)
				if err != nil {
					return nil, fmt.Errorf("error marshaling tool input: %w", err)
				}
				inputJsonString := string(inputJSON)
				result := strings.ReplaceAll(inputJsonString, `"`, "")

				message := fmt.Sprintf("tool: %s %s", req.Name, result)
				pendingMsg := fmt.Sprintf(`{"kind":"tool_call", "message":"%s" , "status": "pending", "operation_id": "%s"}`, message, operationID)
				if err := callback(ctx, pendingMsg); err != nil {
					return nil, fmt.Errorf("error sending pending status: %w", err)
				}
				toolCallsResults += fmt.Sprintf("### Tool call: %s Input: %s:\n", req.Name, inputJsonString)
			}

			log.Printf("⏸️  Operation %s waiting for confirmation...", operationID)

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
				toolCallsResults += fmt.Sprintf("Result: %v\n", output)

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
			fmt.Println("📜 Tools History now has", len(history), "messages")
			fmt.Println(strings.Repeat("-", 20))
		}
	}

	return &toolCallsResult{
		TotalCalls:  totalOfToolsCalls,
		Results:     toolCallsResults,
		LastMessage: lastToolAssistantMessage,
	}, nil
}
