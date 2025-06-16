package cmd

import (
	"github.com/4nkitd/systems-mcp/internal/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "paytring_mcp",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
}

func Execute() {
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		// Handle error
		log.Write("ERROR", err.Error())
	}
}

func init() {

	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringVar(&transport, "transport", "stdio", "Transport type (stdio or sse)")
	rootCmd.PersistentFlags().StringVar(&log_dir, "log_dir", "", "Log directory (default is current directory if not specified)")
	rootCmd.PersistentFlags().StringVar(&host, "host", "localhost", "Host to bind the server to (only for sse transport)")
	rootCmd.PersistentFlags().StringVar(&port, "port", "8080", "Port to bind the server to (only for sse transport)")

	viper.BindPFlag("log_dir", rootCmd.PersistentFlags().Lookup("log_dir"))
	viper.BindPFlag("transport", rootCmd.PersistentFlags().Lookup("transport"))
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))

	rootCmd.AddCommand(serveCmd)

}
