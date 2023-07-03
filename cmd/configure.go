package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const (
	DefaultConfigChoice = "Use default config values (Host: localhost, Port: 9000)"
	SetConfigChoice     = "Set config values now"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure the CLI",
	Long:  `Configure the CLI and set the Host and Port it should use to connect to the server`,
	Run: func(cmd *cobra.Command, args []string) {
		configureCommand()
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func configureCommand() {
	prompt := promptui.Select{
		Label: "Select default config or configure yourself:",
		Items: []string{DefaultConfigChoice, SetConfigChoice},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == DefaultConfigChoice {
		fmt.Println("Default config created.")
		fmt.Println("Host: localhost")
		fmt.Println("Port: 9000")
	}

	if result == SetConfigChoice {
		host, err := hostPrompt()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Host set to %s", host)

		port, err := portPrompt()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Port set to %s", port)
	}
}

func hostPrompt() (string, error) {
	validate := func(input string) error {
		if input == "" {
			return errors.New("Host can't be an empty string")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Host (default: localhost)",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	fmt.Printf("Setting Host to %q\n", result)
	return result, nil
}

func portPrompt() (string, error) {
	validate := func(input string) error {
		port, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid Port number")
		}
		if port < 1023 {
			return errors.New("Please specify a valid port number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Port (default: 9000)",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	fmt.Println("Setting Port to " + result)
	return result, nil
}
