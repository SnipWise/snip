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
	userMessage string,
	//history []*ai.Message,
	operationID string,
	operation *OperationStatus,
	callback core.StreamCallback[string],
) (*toolCallsResult, error) {
	stopped := false
	lastToolAssistantMessage := ""
	totalOfToolsCalls := 0
	toolCallsResults := ""

	history := []*ai.Message{}

	fmt.Println("🟩🟢 MCP 🛠️ Tools", len(config.Tools), "active tools.")
	for _, t := range config.Tools {
		fmt.Println("   -", t.Name())
	}

	fmt.Println("🟧🟠 Initial conversation history:")
	for _, message := range history {
		fmt.Printf("   - History Message - Role: %s, Text: %s\n", message.Role, message.Text())
	}

	for !stopped {
		fmt.Printf("\n🔄 Tool detection loop iteration - Current history length: %d\n", len(history))
		
		resp, err := genkit.Generate(ctx, g,
			ai.WithModelName("openai/"+config.ToolsModel),
			ai.WithSystem(config.ToolsSystemInstruction),
			ai.WithMessages(history...),
			ai.WithPrompt(userMessage),
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
			fmt.Println("✅ No more tool requests, stopping loop")
			break
		}

		fmt.Println("✋ Number of tool requests", len(toolRequests))
		totalOfToolsCalls += len(toolRequests)
		history = append(history, resp.Message)
		fmt.Printf("📥 Added assistant message to history (length now: %d)\n", len(history))

		for _, req := range toolRequests {
			fmt.Println("🛠️ Tool request:", req.Name, "Ref:", req.Ref, "Input:", req.Input)

			// First try to lookup in registered tools (for locally defined tools)
			var tool ai.Tool
			tool = genkit.LookupTool(g, req.Name)
			if tool != nil {
				fmt.Println("   ✅ Found in genkit registry (local tool)")
			}

			// If not found, search in config.Tools (for MCP tools)
			if tool == nil {
				for _, t := range config.Tools {
					if t.Name() == req.Name {
						fmt.Println("   🔍 Found in config.Tools (MCP tool), attempting conversion...")
						// Try to convert ToolRef to Tool
						if toolImpl, ok := t.(ai.Tool); ok {
							tool = toolImpl
							fmt.Println("   ✅ Successfully converted to ai.Tool")
							break
						} else {
							fmt.Println("   ❌ Failed to convert ToolRef to ai.Tool")
						}
					}
				}
			}

			if tool == nil {
				log.Printf("🔴 tool %q not found\n", req.Name)
				//break // [TODO]: continue?
				continue
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
				fmt.Printf("   🔧 Executing tool: %s\n", tool.Name())
				output, err := tool.RunRaw(ctx, req.Input)
				if err != nil {
					log.Fatalf("tool %q execution failed: %v", tool.Name(), err)
				}
				fmt.Println("   🤖 Tool Result:", output)
				toolCallsResults += fmt.Sprintf("Result: %v\n", output)

				part := ai.NewToolResponsePart(&ai.ToolResponse{
					Name:   req.Name,
					Ref:    req.Ref,
					Output: output,
				})
				fmt.Printf("   📝 Adding tool response to history - Name: %s, Ref: %v\n", req.Name, req.Ref)
				history = append(history, ai.NewMessage(ai.RoleTool, nil, part))
				fmt.Printf("   📜 History length now: %d\n", len(history))

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
