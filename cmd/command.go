package cmd

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/manifoldco/promptui"
	"google.golang.org/grpc"

	pb "github.com/danni-popova/analyser/proto"
)

// selectCommand will provide a list of commands for the user to select
// and run
func selectCommand() (string, error) {
	prompt := promptui.Select{
		Label: "What command would you like to run?",
		Items: []string{UserAgentCommand, CookiesCommand, ContentTypeCommand},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}
	return result, nil
}

func runUserAgent(conn *grpc.ClientConn) {
	userAgent, err := userAgentPrompt()
	if err != nil {
		fmt.Println("Error occurred when reading user agent prompt")
		return
	}

	client := pb.NewAnalyserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.AnalyseUserAgent(ctx, &pb.AnalyseUserAgentRequest{
		UserAgent: userAgent,
	})
	if err != nil {
		fmt.Println("Error occurred when analysing user agent string")
		return
	}

	fmt.Println(r.String())
	return
}

func userAgentPrompt() (string, error) {
	validate := func(input string) error {
		if input == "" {
			return errors.New("Host can't be an empty string")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    `User Agent string`,
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return result, nil
}

func runUnimplemented() {
	fmt.Println("Sorry, this command isn't implemented yet.")
}
