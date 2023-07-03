package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// rootCmd represents the base command when called without any subcommands
// Check if the config exists
// Prompt to configure - choose default or configure now
// Check if connection to server is possible with config
// Create config file
// Display list of commands to select - configure, user agent, cookies, etc.
// Run user agent and display results

var rootCmd = &cobra.Command{
	Use:   "analyser",
	Short: "Analyser CLI tool",
	Long:  `Analyser CLI tool`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
