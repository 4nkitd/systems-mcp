package mcp

import (
	"log"

	"github.com/4nkitd/systems-mcp/internal/toolsets"
	"github.com/mark3labs/mcp-go/server"
)

// Config holds the configuration for the MCP server.
type Config struct {
	LogDir      string
	MemoryPath  string
	FetchURLAPI string
}

type ankitd struct {
	LogDir        string
	Mcp           *server.MCPServer
	hooks         *server.Hooks
	memoryTools   *toolsets.MemoryTools
	internetTools *toolsets.InternetTools
}

func New4nkitdMcpServer(config *Config) *ankitd {

	hooks := &server.Hooks{}

	mem, err := toolsets.NewMemory(config.MemoryPath)
	if err != nil {
		log.Fatalf("failed to create memory: %v", err)
	}
	memoryTools := toolsets.NewMemoryTools(mem)

	internetTools := toolsets.NewInternetTools(config.FetchURLAPI)

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
		LogDir:        config.LogDir,
		Mcp:           mcpServer,
		hooks:         hooks,
		memoryTools:   memoryTools,
		internetTools: internetTools,
	}

	// instance.RegisterHooks()
	// instance.RegisterResources()
	// instance.RegisterTools()

	return instance

}
