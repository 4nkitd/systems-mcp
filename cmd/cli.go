package cmd

import (
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

	}
}

func init() {

	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringVar(&key, "key", "", "API key (can be found in your dashboard)")
	rootCmd.PersistentFlags().StringVar(&secret, "secret", "", "API secret (can be found in your dashboard)")
	rootCmd.PersistentFlags().StringVar(&transport, "transport", "stdio", "Transport type (stdio or sse)")
	rootCmd.PersistentFlags().StringVar(&log_dir, "log_dir", "", "Log directory (default is current directory if not specified)")
	rootCmd.MarkPersistentFlagRequired("key")
	rootCmd.MarkPersistentFlagRequired("secret")

	viper.BindPFlag("key", rootCmd.PersistentFlags().Lookup("key"))
	viper.BindPFlag("secret", rootCmd.PersistentFlags().Lookup("secret"))
	viper.BindPFlag("log_dir", rootCmd.PersistentFlags().Lookup("log_dir"))
	viper.BindPFlag("transport", rootCmd.PersistentFlags().Lookup("transport"))

	rootCmd.AddCommand(serveCmd)

}
