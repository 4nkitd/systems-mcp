package toolsets

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

// Order represents an order structure
type Order struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	Items      []Item    `json:"items"`
	Total      float64   `json:"total"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Item represents an item in an order
type Item struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// In-memory storage for orders (in production, this would be a database)
var orders = make(map[string]Order)
var orderCounter = 1

// OrderCreate creates a new order
func OrderCreate(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract required parameters
	customerID, ok := request.Params.Arguments["customer_id"].(string)
	if !ok || customerID == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: "Error: customer_id parameter is required",
				},
			},
			IsError: true,
		}, nil
	}

	// Extract items array
	itemsData, ok := request.Params.Arguments["items"]
	if !ok {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: "Error: items parameter is required",
				},
			},
			IsError: true,
		}, nil
	}

	// Parse items
	var items []Item
	var total float64

	// Handle items as array of maps
	if itemsArray, ok := itemsData.([]interface{}); ok {
		for _, itemData := range itemsArray {
			if itemMap, ok := itemData.(map[string]interface{}); ok {
				item := Item{}

				if name, ok := itemMap["name"].(string); ok {
					item.Name = name
				} else {
					return &mcp.CallToolResult{
						Content: []mcp.Content{
							mcp.TextContent{
								Type: "text",
								Text: "Error: each item must have a name",
							},
						},
						IsError: true,
					}, nil
				}

				if quantity, ok := itemMap["quantity"].(float64); ok {
					item.Quantity = int(quantity)
				} else {
					item.Quantity = 1 // default quantity
				}

				if price, ok := itemMap["price"].(float64); ok {
					item.Price = price
				} else {
					return &mcp.CallToolResult{
						Content: []mcp.Content{
							mcp.TextContent{
								Type: "text",
								Text: "Error: each item must have a price",
							},
						},
						IsError: true,
					}, nil
				}

				item.ID = fmt.Sprintf("item_%d_%d", orderCounter, len(items)+1)
				items = append(items, item)
				total += item.Price * float64(item.Quantity)
			}
		}
	} else {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: "Error: items must be an array of objects",
				},
			},
			IsError: true,
		}, nil
	}

	if len(items) == 0 {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: "Error: at least one item is required",
				},
			},
			IsError: true,
		}, nil
	}

	// Create the order
	orderID := fmt.Sprintf("order_%d", orderCounter)
	orderCounter++

	order := Order{
		ID:         orderID,
		CustomerID: customerID,
		Items:      items,
		Total:      total,
		Status:     "pending",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Store the order
	orders[orderID] = order

	// Return the created order as JSON
	orderJSON, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Error serializing order: %v", err),
				},
			},
			IsError: true,
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: fmt.Sprintf("Order created successfully:\n%s", string(orderJSON)),
			},
		},
	}, nil
}

// OrderFetch fetches order details by ID
func OrderFetch(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract order_id parameter
	orderID, ok := request.Params.Arguments["order_id"].(string)
	if !ok || orderID == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: "Error: order_id parameter is required",
				},
			},
			IsError: true,
		}, nil
	}

	// Fetch the order
	order, exists := orders[orderID]
	if !exists {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Order with ID '%s' not found", orderID),
				},
			},
			IsError: true,
		}, nil
	}

	// Return the order as JSON
	orderJSON, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				mcp.TextContent{
					Type: "text",
					Text: fmt.Sprintf("Error serializing order: %v", err),
				},
			},
			IsError: true,
		}, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: string(orderJSON),
			},
		},
	}, nil
}
