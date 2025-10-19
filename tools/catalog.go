package tools

import (
	"context"
	"fmt"
	"os"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/mcp"
)

type ListToolsInput struct{}

type ToolInfo struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Parameters  map[string]any `json:"parameters,omitempty"`
}

type ListToolsResult struct {
	Tools []ToolInfo `json:"tools"`
	Count int        `json:"count"`
}

// [NOTE]: this is a work-in-progress
func Catalog(ctx context.Context, g *genkit.Genkit, mcpClient *mcp.GenkitMCPClient) []ai.ToolRef {

	toolsList, err := mcpClient.GetActiveTools(ctx, g)

	if err != nil {
		fmt.Println("😡 Error getting the tools list:", err)
		os.Exit(1)
	}
	fmt.Println("🟢 MCP 🛠️ Retrieved", len(toolsList), "active tools from MCP Gateway")

	// Keep MCP tools as ai.Tool (don't convert to ToolRef)
	// This preserves the RunRaw() method needed for execution
	toolRefs := make([]ai.ToolRef, 0, len(toolsList)+1)

	// Add MCP tools directly (ai.Tool implements ai.ToolRef)
	for _, tool := range toolsList {
		toolRefs = append(toolRefs, tool)
	}

	// Define list_tools tool (needs access to toolRefs, so we define it after MCP tools are added)
	listToolsTool := genkit.DefineTool(g, "list_tools", "List all available tools with their descriptions and parameters",
		func(ctx *ai.ToolContext, input ListToolsInput) (ListToolsResult, error) {
			return listTools(toolRefs), nil
		},
	)

	// Append locally defined tools
	toolRefs = append(toolRefs, listToolsTool)

	return toolRefs
}

func listTools(toolRefs []ai.ToolRef) ListToolsResult {
	tools := make([]ToolInfo, 0, len(toolRefs))

	for _, toolRef := range toolRefs {
		toolInfo := ToolInfo{
			Name: toolRef.Name(),
		}

		// Try to get full tool definition if available
		if tool, ok := toolRef.(ai.Tool); ok {
			def := tool.Definition()
			if def != nil {
				toolInfo.Description = def.Description
				toolInfo.Parameters = def.InputSchema
			}
		}

		tools = append(tools, toolInfo)
	}

	return ListToolsResult{
		Tools: tools,
		Count: len(tools),
	}
}
