package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	// prepare payload for AI21 Labs Jurassic 2
	jurassic := bedrock.AI21LabsJurassic{
		Region:            "us-east-1",
		ModelId:           "ai21.j2-mid-v1",
		PromptRequest:     prompt,
		MaxTokensToSample: 200,
		TopP:              1,
		Temperature:       0.7,
		StopSequences:     []string{`""`},
	}

	fmt.Println("Sending prompt to AI21 Jurassic 2")
	
	resp, err = jurassic.InvokeModel()
	if err != nil {
		log.Fatal("error", err)
	}

	text, err = jurassic.GetText(resp)
	if err != nil {
		log.Fatal("error", err)
	}
	fmt.Println(text)

	// prepare payload for Meta Llama
	llama := bedrock.MetaLlama{
		Region:            "us-east-1",
		ModelId:           "meta.llama2-13b-chat-v1",
		Prompt:            prompt,
		MaxTokensToSample: 512,
		TopP:              0.9,
		Temperature:       0.5,
	}

	fmt.Println("Sending prompt to Meta Llama")

	resp, err = llama.InvokeModel()
	if err != nil {
		log.Fatal("error", err)
	}

	text, err = llama.GetText(resp)
	if err != nil {
		log.Fatal("error", err)
	}
	fmt.Println(text)

	// prepare payload for Cohere Command v14
	command := bedrock.CohereCommand{
		Region:            "us-east-1",
		ModelId:           "cohere.command-text-v14",
		Prompt:            prompt,
		Temperature:       0.75,
		TopP:              0.01,
		TopK:              0,
		MaxTokensToSample: 400,
		StopSequences:     []string{`""`},
		ReturnLiklihoods:  "NONE",
		NumGenerations:    1,
	}

	fmt.Println("Sending prompt to Cohere Command")

	resp, err = command.InvokeModel()
	if err != nil {
		log.Fatal("error", err)
	}

	text, err = command.GetText(resp)
	if err != nil {
		log.Fatal("error", err)
	}
	fmt.Println(text)

	// prepare payload for Stability SD

	stability := bedrock.StabilityAISD{
		Region:  "us-east-1",
		ModelId: "stability.stable-diffusion-xl-v0",
		Prompt:  []bedrock.StabilityAISDTextPrompts{{Text: "A cat driving a convertible along the coast"}},
		Scale:   10,
		Seed:    0,
		Steps:   50,
	}

	fmt.Println("Sending prompt to Stability SD")

	resp, err = stability.InvokeModel()
	if err != nil {
		log.Fatal("error", err)
	}

	image, err := stability.GetDecodedImage(resp)
	if err != nil {
		log.Fatal("error", err)
	}
	outputFile := fmt.Sprintf("output-%d.jpg", time.Now().Unix())

	err = os.WriteFile(outputFile, image, 0644)
	if err != nil {
		fmt.Println("error writing to file:", err)
	}

	log.Println("image written to file", outputFile)

}
