package toolsets

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

// Information storage

func SaveInfo(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// save in json and revie from their

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("Information saved under key: %s", "key"),
			},
		},
	}, nil
}

func GetSavedInfo(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	// save in json and revie from their

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: "value",
			},
		},
	}, nil
}
