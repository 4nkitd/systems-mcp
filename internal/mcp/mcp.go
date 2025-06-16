package mcp

import (
	"github.com/mark3labs/mcp-go/server"
)

type ankitd struct {
	LogDir string
	Mcp    *server.MCPServer
	hooks  *server.Hooks
}

func New4nkitdMcpServer(logDir string) *ankitd {

	hooks := &server.Hooks{}

	mcpServer := server.NewMCPServer(
		"4nkitd-mcp-server",
		"0.1.0",
		server.WithResourceCapabilities(false, false),
		server.WithPromptCapabilities(false),
		server.WithToolCapabilities(true),
		server.WithLogging(),
		server.WithHooks(hooks),
	)

	instance := &ankitd{
		LogDir: logDir,
		Mcp:    mcpServer,
		hooks:  hooks,
	}

	// instance.RegisterHooks()
	// instance.RegisterResources()
	// instance.RegisterTools()

	return instance

}
