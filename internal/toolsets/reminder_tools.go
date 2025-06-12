package toolsets

import (
	"context"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func SetAlarm(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract time parameter from request arguments
	timeStr, ok := request.Params.Arguments["time"].(string)
	if !ok || timeStr == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: "Error: time parameter is required in HH:MM format",
				},
			},
			IsError: true,
		}, nil
	}

	// Extract optional message parameter
	message, ok := request.Params.Arguments["message"].(string)
	if !ok || message == "" {
		message = "Alarm!"
	}

	// Parse time in HH:MM format
	alarmTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Invalid time format. Use HH:MM (24-hour format): %v", err),
				},
			},
			IsError: true,
		}, nil
	}

	// Get current time and calculate alarm time for today
	now := time.Now()
	alarmDateTime := time.Date(now.Year(), now.Month(), now.Day(),
		alarmTime.Hour(), alarmTime.Minute(), 0, 0, now.Location())

	// If alarm time has passed today, set it for tomorrow
	if alarmDateTime.Before(now) {
		alarmDateTime = alarmDateTime.Add(24 * time.Hour)
	}

	formattedAlarmTime := alarmDateTime.Format("2006-01-02 15:04:05")

	// Start alarm in a goroutine
	go func(at time.Time, msg string) {
		duration := time.Until(at)
		time.Sleep(duration)

		// Simple alarm notification - could be enhanced with system notifications
		fmt.Printf("\n🔔 ALARM: %s - %s\n", at.Format("15:04:05"), msg)
	}(alarmDateTime, message)

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("Alarm set for %s with message: %s", formattedAlarmTime, message),
			},
		},
	}, nil
}
