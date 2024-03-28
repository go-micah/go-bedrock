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

	prompt := "Please write me a short poem about a chicken."

	// prepare payload for Mixtral Instruct
	body := providers.MistralMixtralInstructInvokeModelInput{
		Prompt:            prompt,
		MaxTokensToSample: 500,
		TopP:              0.9,
		TopK:              200,
		Temperature:       0.5,
		StopSequences:     []string{},
	}

	fmt.Println("Sending prompt to Mixtral Instruct")
	fmt.Println("")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}

	svc := bedrockruntime.NewFromConfig(cfg)

	accept := "*/*"
	contentType := "application/json"
	modelId := "mistral.mistral-7b-instruct-v0:2"

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

	var out providers.MistralMixtralInstructInvokeModelOutputs

	err = json.Unmarshal(resp.Body, &out)
	if err != nil {
		fmt.Printf("unable to Unmarshal JSON, %v", err)
	}

	fmt.Println(out.Output[0].Text)

}
