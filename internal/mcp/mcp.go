package mcp

import (
	"github.com/mark3labs/mcp-go/server"
)

type Paytring struct {
	LogDir string
	Mcp    *server.MCPServer
	hooks  *server.Hooks
}

func NewPaytringMcpServer(logDir string) *Paytring {

	hooks := &server.Hooks{}

	mcpServer := server.NewMCPServer(
		"paytring-mcp-server",
		"0.1.0",
		server.WithResourceCapabilities(false, false),
		server.WithPromptCapabilities(false),
		server.WithToolCapabilities(true),
		server.WithLogging(),
		server.WithHooks(hooks),
	)

	instance := &Paytring{
		LogDir: logDir,
		Mcp:    mcpServer,
		hooks:  hooks,
	}

	// instance.RegisterHooks()
	// instance.RegisterResources()
	// instance.RegisterTools()

	return instance

}
