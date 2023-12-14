package main

import (
	"fmt"
	"log"

	"github.com/go-micah/go-bedrock"
)

func main() {

	prompt := "Please write me a short poem about a chicken"

	// prepare payload for Anthropic Claude v2
	claude := bedrock.AnthropicClaude{
		Region:            "us-east-1",
		ModelId:           "anthropic.claude-v2",
		Prompt:            "Human: \n\nHuman: " + prompt + "\n\nAssistant:",
		MaxTokensToSample: 500,
		TopP:              0.999,
		TopK:              250,
		Temperature:       1,
		StopSequences:     []string{`"\n\nHuman:\"`},
	}

	fmt.Println("Sending prompt to Anthropic Claude v2")

	resp, err := claude.InvokeModel()
	if err != nil {
		log.Fatal("error", err)
	}

	text, err := claude.GetText(resp)
	if err != nil {
		log.Fatal("error", err)
	}
	fmt.Println(text)

}
