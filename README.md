[![Go Reference](https://pkg.go.dev/badge/github.com/go-micah/go-bedrock.svg)](https://pkg.go.dev/github.com/go-micah/go-bedrock)

# go-bedrock

A wrapper around the Amazon Bedrock API written in Go

## Use

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/go-micah/go-bedrock/providers"
)

func main() {

	prompt := "Please write me a short poem about a chicken"

	// prepare payload for Anthropic Claude v2
	body := providers.AnthropicClaudeInvokeModelInput{
		Prompt:            "Human: \n\nHuman: " + prompt + "\n\nAssistant:",
		MaxTokensToSample: 500,
		TopP:              0.999,
		TopK:              250,
		Temperature:       1,
		StopSequences:     []string{`"\n\nHuman:\"`},
	}

	fmt.Println("Sending prompt to Anthropic Claude v2")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}

	svc := bedrockruntime.NewFromConfig(cfg)

	accept := "*/*"
	contentType := "application/json"
	modelId := "anthropic.claude-v2"

	bodyString, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("unable to Marshal, %v", err)
	}

	resp, err := svc.InvokeModel(context.TODO(), &bedrockruntime.InvokeModelInput{
		Accept:      &accept,
		ModelId:     &modelId,
		ContentType: &contentType,
		Body:        bodyString,
	})
	if err != nil {
		log.Fatalf("error from Bedrock, %v", err)
	}

	var out providers.AnthropicClaudeInvokeModelOutput

	err = json.Unmarshal(resp.Body, &out)
	if err != nil {
		fmt.Printf("unable to Unmarshal JSON, %v", err)
	}

	fmt.Println(out.Completion)

}

```
