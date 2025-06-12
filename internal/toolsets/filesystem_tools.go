package toolsets

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

func GetCurrentWorkingDirectory(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Error getting current directory: %v", err),
				},
			},
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: cwd,
			},
		},
	}, nil
}

func ListDirectory(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	if path == "" {
		var err error
		path, err = os.Getwd()
		if err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: fmt.Sprintf("Error getting current directory: %v", err),
					},
				},
			}, nil
		}
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Error reading directory %s: %v", path, err),
				},
			},
		}, nil
	}

	var result strings.Builder
	result.WriteString(fmt.Sprintf("Contents of %s:\n", path))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue // Skip entries we can't get info for
		}

		var entryType string
		if entry.IsDir() {
			entryType = "DIR"
		} else {
			entryType = "FILE"
		}

		result.WriteString(fmt.Sprintf("%-5s %10d %s %s\n",
			entryType,
			info.Size(),
			info.ModTime().Format("2006-01-02 15:04:05"),
			entry.Name()))
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: result.String(),
			},
		},
	}, nil
}

func ReadFile(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Error reading file %s: %v", path, err),
				},
			},
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: string(content),
			},
		},
	}, nil
}
