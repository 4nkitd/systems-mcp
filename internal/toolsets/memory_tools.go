package toolsets

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

// MemoryTools provides tool functions for memory operations.
type MemoryTools struct {
	mem *Memory
}

// NewMemoryTools creates a new MemoryTools instance.
func NewMemoryTools(mem *Memory) *MemoryTools {
	return &MemoryTools{mem: mem}
}

// SaveInfo saves a key-value pair to the memory file.
func (t *MemoryTools) SaveInfo(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

	if err := t.mem.Set(key, value); err != nil {
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

// GetSavedInfo retrieves information from the memory file.
func (t *MemoryTools) GetSavedInfo(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.Params.Arguments
	if keyVal, ok := args["key"].(string); ok && keyVal != "" {
		val, err := t.mem.Get(keyVal)
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{Type: "text", Text: err.Error()},
				},
				IsError: true,
			}, nil
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{Type: "text", Text: val},
			},
		}, nil
	}

	// No key provided, return all stored data
	all, err := t.mem.GetAll()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{Type: "text", Text: fmt.Sprintf("Error getting all data: %v", err)},
			},
			IsError: true,
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{Type: "text", Text: fmt.Sprintf("Stored information:\n%s", all)},
		},
	}, nil
}
