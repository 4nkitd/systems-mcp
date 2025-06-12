package toolsets

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/mark3labs/mcp-go/mcp"
)

func VolumeUp(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("osascript", "-e", "set volume output volume (output volume of (get volume settings) + 10)")
	case "linux":
		// Ensure amixer (from alsa-utils) is installed
		if _, err := exec.LookPath("amixer"); err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: "amixer is not installed. Please install alsa-utils.",
					},
				},
			}, nil
		}
		cmd = exec.Command("amixer", "sset", "Master", "5%+")
	case "windows":
		// This PowerShell command sends the Volume Up key press
		cmd = exec.Command("powershell", "-Command", "(New-Object -ComObject WScript.Shell).SendKeys([char]175)")
	default:
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: "Volume control not supported on this platform",
				},
			},
		}, nil
	}

	err := cmd.Run()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Error increasing volume: %v. Command: %s", err, cmd.String()),
				},
			},
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: "Volume increased",
			},
		},
	}, nil
}

func VolumeDown(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("osascript", "-e", "set volume output volume (output volume of (get volume settings) - 10)")
	case "linux":
		if _, err := exec.LookPath("amixer"); err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: "amixer is not installed. Please install alsa-utils.",
					},
				},
			}, nil
		}
		cmd = exec.Command("amixer", "sset", "Master", "5%-")
	case "windows":
		cmd = exec.Command("powershell", "-Command", "(New-Object -ComObject WScript.Shell).SendKeys([char]174)")
	default:
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: "Volume control not supported on this platform",
				},
			},
		}, nil
	}

	err := cmd.Run()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Error decreasing volume: %v. Command: %s", err, cmd.String()),
				},
			},
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: "Volume decreased",
			},
		},
	}, nil
}

func VolumeMute(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("osascript", "-e", "set volume with output muted")
	case "linux":
		if _, err := exec.LookPath("amixer"); err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: "amixer is not installed. Please install alsa-utils.",
					},
				},
			}, nil
		}
		cmd = exec.Command("amixer", "sset", "Master", "mute")
	case "windows":
		cmd = exec.Command("powershell", "-Command", "(New-Object -ComObject WScript.Shell).SendKeys([char]173)")
	default:
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: "Volume control not supported on this platform",
				},
			},
		}, nil
	}

	err := cmd.Run()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Error muting volume: %v. Command: %s", err, cmd.String()),
				},
			},
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: "Volume muted",
			},
		},
	}, nil
}

func VolumeUnmute(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("osascript", "-e", "set volume without output muted")
	case "linux":
		if _, err := exec.LookPath("amixer"); err != nil {
			return &mcp.CallToolResult{
				Content: []mcp.Content{
					mcp.TextContent{
						Type: "text",
						Text: "amixer is not installed. Please install alsa-utils.",
					},
				},
			}, nil
		}
		cmd = exec.Command("amixer", "sset", "Master", "unmute")
	case "windows":
		// Toggling mute often uses the same key, so this might work for unmute as well.
		// If a separate unmute command is needed, it would be platform-specific.
		cmd = exec.Command("powershell", "-Command", "(New-Object -ComObject WScript.Shell).SendKeys([char]173)")
	default:
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: "Volume control not supported on this platform",
				},
			},
		}, nil
	}

	err := cmd.Run()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Error unmuting volume: %v. Command: %s", err, cmd.String()),
				},
			},
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: "Volume unmuted",
			},
		},
	}, nil
}

func Speak(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var cmd *exec.Cmd

	cmd = exec.Command("say", message)

	err := cmd.Run()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Error increasing volume: %v. Command: %s", err, cmd.String()),
				},
			},
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: "Volume increased",
			},
		},
	}, nil
}
