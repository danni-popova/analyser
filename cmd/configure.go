package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

const (
	DefaultConfigChoice = "Use default config values (Host: localhost, Port: 9000)"
	SetConfigChoice     = "Set config values now"
)

// runConfigure will prompt the user to create a new config
func runConfigure() (*Config, error) {
	var config *Config

	prompt := promptui.Select{
		Label: "Select default config or configure yourself:",
		Items: []string{DefaultConfigChoice, SetConfigChoice},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	switch result {
	case DefaultConfigChoice:
		fmt.Println("Using default config")
		config = &Config{
			Host: "localhost",
			Port: 9000,
		}
	case SetConfigChoice:
		host, err := hostPrompt()
		if err != nil {
			// TODO: remove
			fmt.Println(err.Error())
			return nil, errors.New("couldn't configure cli")
		}
		port, err := portPrompt()
		if err != nil {
			// TODO: remove
			fmt.Println(err.Error())
			return nil, errors.New("couldn't configure cli")
		}

		portInt, _ := strconv.ParseFloat(port, 64)
		config = &Config{
			Host: host,
			Port: int(portInt),
		}
	}

	return config, nil
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
	// TODO: change this to return number
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
