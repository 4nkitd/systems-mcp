package mcp

import (
	"context"
	"fmt"

	"github.com/4nkitd/systems-mcp/internal/log"
	"github.com/mark3labs/mcp-go/mcp"
)

func (p *4nkitd) RegisterHooks() {

	p.hooks.AddBeforeAny(func(ctx context.Context, id any, method mcp.MCPMethod, message any) {
		log.Write("DEBUG", fmt.Sprintf("beforeAny: %s, %v, %v\n", method, id, message))
	})
	p.hooks.AddOnSuccess(func(ctx context.Context, id any, method mcp.MCPMethod, message any, result any) {
		log.Write("DEBUG", fmt.Sprintf("onSuccess: %s, %v, %v, %v\n", method, id, message, result))
	})
	p.hooks.AddOnError(func(ctx context.Context, id any, method mcp.MCPMethod, message any, err error) {
		log.Write("DEBUG", fmt.Sprintf("onError: %s, %v, %v, %v\n", method, id, message, err))
	})
	p.hooks.AddBeforeInitialize(func(ctx context.Context, id any, message *mcp.InitializeRequest) {
		log.Write("DEBUG", fmt.Sprintf("beforeInitialize: %v, %v\n", id, message))
	})
	p.hooks.AddAfterInitialize(func(ctx context.Context, id any, message *mcp.InitializeRequest, result *mcp.InitializeResult) {
		log.Write("DEBUG", fmt.Sprintf("afterInitialize: %v, %v, %v\n", id, message, result))
	})
	p.hooks.AddAfterCallTool(func(ctx context.Context, id any, message *mcp.CallToolRequest, result *mcp.CallToolResult) {
		log.Write("DEBUG", fmt.Sprintf("afterCallTool: %v, %v, %v\n", id, message, result))
	})
	p.hooks.AddBeforeCallTool(func(ctx context.Context, id any, message *mcp.CallToolRequest) {
		log.Write("DEBUG", fmt.Sprintf("beforeCallTool: %v, %v\n", id, message))
	})

}
