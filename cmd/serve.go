package cmd

import (
	"log"

	"github.com/4nkitd/systems-mcp/internal/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the 4nkitd mcp server",
	Long:  `This command starts the 4nkitd mcp server. It will listen for incoming requests and process them accordingly.`,
	Run:   serveServer,
}

func serveServer(cmd *cobra.Command, args []string) {

	log.Println("Starting the 4nkitd mcp server...")

	transport := viper.GetString("transport")
	host := viper.GetString("host")
	port := viper.GetString("port")

	log.Println("Log Directory:", viper.GetString("log_dir"))
	log.Println("Transport:", transport)

	mcpServer := mcp.New4nkitdMcpServer(viper.GetString("log_dir"))
	mcpServer.RegisterHooks()
	mcpServer.RegisterTools()

	// Only check for "sse" since stdio is the default
	if transport == "sse" {
		address := host + ":" + port
		baseURL := "http://" + address

		sseServer := server.NewSSEServer(mcpServer.Mcp, server.WithBaseURL(baseURL))

		log.Println("========================================")
		log.Printf("🚀 MCP Server running!")
		log.Printf("📍 URL: %s", baseURL)
		log.Printf("🔧 Transport: SSE")
		log.Printf("📂 Log Directory: %s", viper.GetString("log_dir"))
		log.Println("========================================")

		if err := sseServer.Start(":" + port); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	} else {
		log.Println("========================================")
		log.Printf("🚀 MCP Server running!")
		log.Printf("🔧 Transport: STDIO")
		log.Printf("📂 Log Directory: %s", viper.GetString("log_dir"))
		log.Println("ℹ️  Using standard input/output for communication")
		log.Println("========================================")

		if err := server.ServeStdio(mcpServer.Mcp); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}

}
