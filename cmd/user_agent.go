package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/danni-popova/analyser/proto"
)

// userAgentCmd represents the userAgent command
var userAgentCmd = &cobra.Command{
	Use:   "user-agent",
	Short: "Analyse user agent string",
	Long: `Analyse user agent string

The user agent command accepts an HTTP user agent string,
analyses it and returns information about it.

Example user agent: "Mozilla/5.0 (Android 4.3; Mobile; rv:54.0) Gecko/54.0 Firefox/54.0"
`,
	Run: runUserAgentCommand,
}

func init() {
	rootCmd.AddCommand(userAgentCmd)

	// Here you will define your flags and configuration settings.
	readConfig()
}

func runUserAgentCommand(cmd *cobra.Command, args []string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Couldn't connect to server. Check you have the right values for Host and Port in your config.")
	}
	defer conn.Close()

	client := pb.NewAnalyserClient(conn)
	ctx, cancel := context.WithTimeout(cmd.Context(), time.Second*5)
	defer cancel()

	r, err := client.AnalyseUserAgent(ctx, &pb.AnalyseUserAgentRequest{
		UserAgent: "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println("Browser: " + r.Browser)
}

func readConfig() {
	var config Config
	configFile, err := os.Open("config.json")
	// If the file doesn't exist, we create a default one
	if err != nil {
		log.Println("Couldn't find default config")
		writeDefaultConfig()
	}

	// Ensure the config is valid and can be parsed
	// Otherwise overwrite with default
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		log.Println("Couldn't parse config")
		writeDefaultConfig()
	}
}

func writeDefaultConfig() {
	log.Println("Creating default config")
	config := &Config{
		Host: "localhost",
		Port: 9000,
	}

	configContents, err := json.Marshal(config)
	if err != nil {
		log.Fatal("couldn't marshal default config")
	}

	err = os.WriteFile("config.json", configContents, 0644)
	if err != nil {
		log.Fatal("couldn't write default config")
	}
}
