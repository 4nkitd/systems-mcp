package cmd

import (
	"log"

	"github.com/4nkitd/mcp/internal/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the paytring mcp server",
	Long:  `This command starts the paytring mcp server. It will listen for incoming requests and process them accordingly.`,
	Run:   serveServer,
}

func serveServer(cmd *cobra.Command, args []string) {

	log.Println("Starting the paytring mcp server...")

	key := viper.GetString("key")
	secret := viper.GetString("secret")
	transport := viper.GetString("transport")

	log.Println("Key:", key)
	log.Println("Secret:", secret)
	log.Println("Log Directory:", viper.GetString("log_dir"))

	mcpServer := mcp.NewPaytringMcpServer(key, secret, viper.GetString("log_dir"))
	mcpServer.RegisterHooks()
	mcpServer.RegisterTools()

	// Only check for "sse" since stdio is the default
	if transport == "sse" {
		sseServer := server.NewSSEServer(mcpServer.Mcp, server.WithBaseURL("http://localhost:8080"))
		log.Printf("SSE server listening on :8080")
		if err := sseServer.Start(":8080"); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	} else {
		if err := server.ServeStdio(mcpServer.Mcp); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}

}
