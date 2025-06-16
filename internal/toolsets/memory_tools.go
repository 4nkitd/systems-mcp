package toolsets

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
)

// Information storage

func SaveInfo(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Define storage file
	const storageFile = "memory_store.json"

	// Extract key and value from arguments
	args := request.Params.Arguments
	key, ok := args["key"].(string)
	if !ok || key == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{Type: "text", Text: "Error: 'key' parameter is required"},
			},
			IsError: true,
		}, nil
	}
	value, ok := args["value"].(string)
	if !ok {
		// Try to marshal non-string value
		raw, _ := json.Marshal(args["value"])
		value = string(raw)
	}

	// Load existing data
	data := make(map[string]string)
	if b, err := os.ReadFile(storageFile); err == nil {
		_ = json.Unmarshal(b, &data)
	}
	// Update and save
	data[key] = value
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{Type: "text", Text: fmt.Sprintf("Error marshaling data: %v", err)},
			},
			IsError: true,
		}, nil
	}
	if err := os.WriteFile(storageFile, b, 0644); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{Type: "text", Text: fmt.Sprintf("Error saving data: %v", err)},
			},
			IsError: true,
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{Type: "text", Text: fmt.Sprintf("Information saved under key: %s", key)},
		},
	}, nil
}

func GetSavedInfo(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Define storage file
	const storageFile = "memory_store.json"

	// Load existing data
	data := make(map[string]string)
	b, err := os.ReadFile(storageFile)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{Type: "text", Text: fmt.Sprintf("Error reading storage: %v", err)},
			},
			IsError: true,
		}, nil
	}
	_ = json.Unmarshal(b, &data)

	// Extract key from arguments
	args := request.Params.Arguments
	if keyVal, ok := args["key"].(string); ok && keyVal != "" {
		if val, exists := data[keyVal]; exists {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{Type: "text", Text: val},
				},
			}, nil
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{Type: "text", Text: fmt.Sprintf("No information found for key: %s", keyVal)},
			},
			IsError: true,
		}, nil
	}

	// No key provided, return all stored data
	all, _ := json.MarshalIndent(data, "", "  ")
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{Type: "text", Text: fmt.Sprintf("Stored information:\n%s", string(all))},
		},
	}, nil
}
