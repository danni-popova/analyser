package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	UserAgentCommand = "user-agent"
)

var cfgFile string

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "analyser",
	Short: "Analyser CLI tool",
	Long:  `Analyser CLI tool`,
	Run: func(cmd *cobra.Command, args []string) {
		config := &Config{}

		// Check if the config exists
		config, err := readConfig()

		// If there was an error, prompt the user to configure
		if err != nil {
			// Prompt to configure - choose default or configure now
			fmt.Println("Welcome to the Analyser CLI! We are going to run you through the steps to configure the tool.")
			config, err = runConfigure()
			if err != nil {
				fmt.Printf("couldn't configure CLI: %v\n", err)
				return
			}
		}

		// Check if the server is reachable with the provided config
		err = testConnection(config)
		if err != nil {
			fmt.Println("Couldn't connect to server. Please check you configured the correct values.")
			return
		}

		// Create config file
		err = createConfig(config)
		if err != nil {
			fmt.Println("Oops! Couldn't persist the config.")
		}

		command, err := selectCommand()
		if err != nil {
			fmt.Println("Oops! Error selecting command occurred.")
			return
		}
		// Run user agent and display results
		if command == UserAgentCommand {
			runUserAgent(config)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// readConfig will detect if there's an existing configuration
func readConfig() (*Config, error) {

	return nil, errors.New("config file not found")
}

// runConfigure will prompt the user to create a new config
func runConfigure() (*Config, error) {

	return nil, errors.New("couldn't configure cli")
}

// testConnection will ensure the server is reachable
func testConnection(config *Config) error {
	return nil
}

// createConfig will persist the config so the user doesn't need to
// reconfigure on next run
func createConfig(config *Config) error {
	return nil
}

// selectCommand will provide a list of commands for the user to select
// and run
func selectCommand() (string, error) {
	return "", errors.New("error occurred while prompting to select command to run")
}

func runUserAgent(config *Config) {

	return
}

func runUnimplemented() {
	fmt.Println("Sorry, this command isn't implemented yet.")
}
