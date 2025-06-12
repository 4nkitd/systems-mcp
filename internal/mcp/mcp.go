package mcp

import (
	"os"

	"github.com/mark3labs/mcp-go/server"
)

type Paytring struct {
	Key    string
	Secret string
	LogDir string
	Mcp    *server.MCPServer
	hooks  *server.Hooks
}

func NewPaytringMcpServer(key, secret, logDir string) *Paytring {

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

	os.Setenv("PAYTRING_API_KEY", key)
	os.Setenv("PAYTRING_API_SECRET", secret)

	instance := &Paytring{
		Key:    key,
		Secret: secret,
		LogDir: logDir,
		Mcp:    mcpServer,
		hooks:  hooks,
	}

	// instance.RegisterHooks()
	// instance.RegisterResources()
	// instance.RegisterTools()

	return instance

}
