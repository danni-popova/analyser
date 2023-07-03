package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	UserAgentCommand   = "user-agent"
	CookiesCommand     = "cookies"
	ContentTypeCommand = "content-type"
)

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
		conn, err := testConnection(config)
		if err != nil {
			fmt.Println("Couldn't connect to server. Please check you configured the correct values.")
			return
		}
		// Make sure to close connection on exit
		defer conn.Close()

		// Create config file
		err = createConfig(config)
		// TODO: this error means the config already existed, this should maybe be moved
		if err != nil {
			fmt.Println("Oops! Couldn't persist the config.")
		}

		command, err := selectCommand()
		if err != nil {
			fmt.Println("Oops! Error selecting command occurred.")
			return
		}
		// Run user agent and display results
		switch command {
		case UserAgentCommand:
			runUserAgent(conn)
		default:
			runUnimplemented()
		}

		// Display options again
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
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name ".analyser" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".analyser")
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		return nil, errors.New("config file not found")
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
	host := viper.GetString("host")
	if err != nil {
		return nil, errors.New("Host isn't set in config")
	}

	port := viper.GetInt("port")
	if err != nil {
		return nil, errors.New("Port isn't set in config")
	}

	return &Config{
		Host: host,
		Port: port,
	}, nil
}

// testConnection will ensure the server is reachable
func testConnection(config *Config) (*grpc.ClientConn, error) {
	connString := fmt.Sprintf("%s:%d", config.Host, config.Port)
	return grpc.Dial(connString, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

// createConfig will persist the config so the user doesn't need to
// reconfigure on next run
func createConfig(config *Config) error {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".analyser")
	viper.SetConfigType("json")
	viper.Set("host", config.Host)
	viper.Set("port", config.Port)

	err = viper.SafeWriteConfig()

	if err != nil {
		return err
	}

	return nil
}
