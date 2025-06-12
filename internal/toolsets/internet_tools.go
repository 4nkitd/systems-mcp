package toolsets

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

func GetWeather(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract location parameter from request arguments
	location, ok := request.Params.Arguments["location"].(string)
	if !ok || location == "" {
		// Attempt to get current location for weather if not specified
		locResult, err := GetCurrentLocation(ctx, request)
		if err == nil && len(locResult.Content) > 0 {
			if textContent, ok := locResult.Content[0].(mcp.TextContent); ok {
				// Crude parsing of city from getCurrentLocation output for wttr.in
				// This is highly dependent on ipinfo.io output format
				lines := strings.Split(textContent.Text, "\n")
				for _, line := range lines {
					if strings.HasPrefix(line, "City:") {
						location = strings.TrimSpace(strings.TrimPrefix(line, "City:"))
						break
					}
				}
			}
		}
		if location == "" { // Default if location couldn't be determined
			location = "current" // wttr.in understands "current"
		}
	}

	// Using wttr.in service as an example
	// Ensure curl is installed
	if _, err := exec.LookPath("curl"); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: "curl is not installed. Please install curl to fetch weather information.",
				},
			},
		}, nil
	}

	cmd := exec.Command("curl", "-s", fmt.Sprintf("http://wttr.in/%s?format=3", location))
	output, err := cmd.Output()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Weather service unavailable or error fetching weather for '%s': %v. Command: %s", location, err, cmd.String()),
				},
			},
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: strings.TrimSpace(string(output)),
			},
		},
	}, nil
}

func GetCurrentLocation(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Ensure curl is installed
	if _, err := exec.LookPath("curl"); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: "curl is not installed. Please install curl to fetch location information.",
				},
			},
		}, nil
	}

	// Using ipinfo.io service
	cmd := exec.Command("curl", "-s", "http://ipinfo.io/json")
	output, err := cmd.Output()
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Location service unavailable: %v. Command: %s", err, cmd.String()),
				},
			},
		}, nil
	}

	var locationData map[string]interface{}
	if err := json.Unmarshal(output, &locationData); err != nil {
		// If unmarshal fails, return the raw output as it might contain an error message from the service
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Error parsing location data: %v. Raw response: %s", err, string(output)),
				},
			},
		}, nil
	}

	// Format the location data
	var result strings.Builder
	if ip, ok := locationData["ip"].(string); ok {
		result.WriteString(fmt.Sprintf("IP: %s\n", ip))
	}
	if city, ok := locationData["city"].(string); ok {
		result.WriteString(fmt.Sprintf("City: %s\n", city))
	}
	if region, ok := locationData["region"].(string); ok {
		result.WriteString(fmt.Sprintf("Region: %s\n", region))
	}
	if country, ok := locationData["country"].(string); ok {
		result.WriteString(fmt.Sprintf("Country: %s\n", country))
	}
	if loc, ok := locationData["loc"].(string); ok {
		result.WriteString(fmt.Sprintf("Coordinates: %s\n", loc))
	}
	if org, ok := locationData["org"].(string); ok {
		result.WriteString(fmt.Sprintf("Organization: %s\n", org))
	}
	if postal, ok := locationData["postal"].(string); ok {
		result.WriteString(fmt.Sprintf("Postal Code: %s\n", postal))
	}
	if timezone, ok := locationData["timezone"].(string); ok {
		result.WriteString(fmt.Sprintf("Timezone: %s\n", timezone))
	}

	if result.Len() == 0 {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("No location details found in response: %s", string(output)),
				},
			},
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: strings.TrimSpace(result.String()),
			},
		},
	}, nil
}
