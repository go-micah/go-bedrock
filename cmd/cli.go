package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-micah/go-bedrock"
)

func main() {

	prompt := "Please write me a short poem about a chicken"

	var options bedrock.Options

	options.ModelID = "anthropic.claude-v2"
	options.Region = "us-east-1"

	options.MaxTokensToSample = 500

	options.Temperature = 1
	options.TopP = 0.999
	options.TopK = 250

	options.StopSequences = []string{`"\n\nHuman:\"`}

	resp, err := bedrock.SendToBedrock(prompt, options)
	if err != nil {
		log.Fatal("error", err)
	}

	var response bedrock.AnthropicResponse

	err = json.Unmarshal(resp.Body, &response)

	if err != nil {
		log.Fatal("failed to unmarshal", err)
	}

	fmt.Print(response.Completion)
}
