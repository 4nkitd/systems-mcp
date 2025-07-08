package toolsets

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/mark3labs/mcp-go/mcp"
)

// InternetTools provides tool functions for internet operations.
type InternetTools struct {
	FetchURLAPI string
}

// NewInternetTools creates a new InternetTools instance.
func NewInternetTools(fetchURLAPI string) *InternetTools {
	if fetchURLAPI == "" {
		fetchURLAPI = "https://md.dhr.wtf/"
	}
	return &InternetTools{FetchURLAPI: fetchURLAPI}
}

func (t *InternetTools) GetWeather(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract location parameter from request arguments
	location, ok := request.Params.Arguments["location"].(string)
	if !ok || location == "" {
		// Attempt to get current location for weather if not specified
		locResult, err := t.GetCurrentLocation(ctx, request)
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
	client := resty.New()
	resp, err := client.R().Get(fmt.Sprintf("http://wttr.in/%s?format=3", location))

	if err != nil || resp.IsError() {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Weather service unavailable or error fetching weather for '%s': %v", location, err),
				},
			},
			IsError: true,
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: strings.TrimSpace(resp.String()),
			},
		},
	}, nil
}

func (t *InternetTools) GetCurrentLocation(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Using ipinfo.io service
	client := resty.New()
	resp, err := client.R().Get("http://ipinfo.io/json")

	if err != nil || resp.IsError() {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Location service unavailable: %v", err),
				},
			},
			IsError: true,
		}, nil
	}

	var locationData map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &locationData); err != nil {
		// If unmarshal fails, return the raw output as it might contain an error message from the service
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Error parsing location data: %v. Raw response: %s", err, resp.String()),
				},
			},
			IsError: true,
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
					Text: fmt.Sprintf("No location details found in response: %s", resp.String()),
				},
			},
			IsError: true,
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

func (t *InternetTools) FetchURL(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract url parameter from request arguments
	url, ok := request.Params.Arguments["url"].(string)
	if !ok || url == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{Type: "text", Text: "Error: 'url' parameter is required"},
			},
			IsError: true,
		}, nil
	}

	// Using md.dhr.wtf service
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("url", url).
		Get(t.FetchURLAPI)

	if err != nil || resp.IsError() {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Error fetching URL '%s': %v", url, err),
				},
			},
			IsError: true,
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: strings.TrimSpace(resp.String()),
			},
		},
	}, nil
}
