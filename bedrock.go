// Package bedrock is a wrapper around the Amazon Bedrock API
package bedrock

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

// Options is a struct that represents all the available LLM options and prompt data
type Options struct {
	Document          string
	Region            string
	ModelID           string
	MaxTokensToSample int
	TopP              float64
	TopK              int
	Temperature       float64
	StopSequences     []string
	ReturnLiklihoods  string
	Stream            bool
	Scale             int
	Seed              int
	Steps             int
}

// AnthropicResponse is a struct that represents the response from Bedrock
type AnthropicResponse struct {
	Completion string
}

// CohereResponse is a struct that represents the response from Cohere
type CohereResponse struct {
	Generations []CohereResponseGeneration `json:"generations"`
}

// CohereResponseGeneration is a struct that represents a generation from Cohere
type CohereResponseGeneration struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// AI21Response is a struct that represents the response from Cohere
type AI21Response struct {
	Completion string
}

// MetaResponse is a struct that represents the response from Bedrock
type MetaResponse struct {
	Generation string `json:"generation"`
}

// StabilityResponse is a struct that represents the response from Stability
type StabilityResponse struct {
	Result    string              `json:"result"`
	Artifacts []StabilityArtifact `json:"artifacts"`
}

// StabilityArtifact is a struct that represents an artifact from Stability
type StabilityArtifact struct {
	Base64       string `json:"base64"`
	FinishReason string `json:"finishReason"`
}

// AnthropicPayloadBody is a struct that represents the payload body for the post request to Bedrock
type AnthropicPayloadBody struct {
	Prompt            string   `json:"prompt"`
	MaxTokensToSample int      `json:"max_tokens_to_sample"`
	Temperature       float64  `json:"temperature"`
	TopK              int      `json:"top_k"`
	TopP              float64  `json:"top_p"`
	StopSequences     []string `json:"stop_sequences"`
}

// A121PayloadBody is a struct that represents the payload body for the post request to Bedrock
type AI21PayloadBody struct {
	Prompt            string   `json:"prompt"`
	Temperature       float64  `json:"temperature"`
	TopP              float64  `json:"topP"`
	MaxTokensToSample int      `json:"maxTokens"`
	StopSequences     []string `json:"stopSequences"`
}

// CoherePayloadBody is a struct that represents the payload body for the post request to Bedrock
type CoherePayloadBody struct {
	Prompt            string   `json:"prompt"`
	Temperature       float64  `json:"temperature"`
	TopP              float64  `json:"p"`
	TopK              float64  `json:"k"`
	MaxTokensToSample int      `json:"max_tokens"`
	StopSequences     []string `json:"stop_sequences"`
	ReturnLiklihoods  string   `json:"return_likelihoods"`
	Stream            bool     `json:"stream"`
	// Generations       int      `json:"num_generations"`
}

// MetaPayloadBody is a struct that represents the payload body for the post request to Bedrock
type MetaPayloadBody struct {
	Prompt            string  `json:"prompt"`
	Temperature       float64 `json:"temperature"`
	TopP              float64 `json:"top_p"`
	MaxTokensToSample int     `json:"max_gen_len"`
}

// StabilityTextPrompts is a struct that represents the text prompts for the post request to Bedrock
type StabilityTextPrompts struct {
	Text string `json:"text"`
}

// StabilityPayloadBody is a struct that represents the payload body for the post request to Bedrock
type StabilityPayloadBody struct {
	Prompt []StabilityTextPrompts `json:"text_prompts"`
	Scale  float64                `json:"cfg_scale"`
	Steps  int                    `json:"steps"`
	Seed   int                    `json:"seed"`
}

// DecodeImage is a function that decodes the image from the response
func (a *StabilityArtifact) DecodeImage() ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(a.Base64)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

// SerializePayload is a function that serializes the payload body before sending to Bedrock
func SerializePayload(prompt string, options Options) ([]byte, error) {

	if options.Document != "" {
		prompt = options.Document + prompt
	}

	model := options.ModelID
	modelTLD := model[:strings.IndexByte(model, '.')]

	// if config says anthropic, use AnthropicPayloadBody
	if modelTLD == "anthropic" {

		var body AnthropicPayloadBody
		body.Prompt = "Human: \n\nHuman: " + prompt + "\n\nAssistant:"
		body.MaxTokensToSample = options.MaxTokensToSample
		body.Temperature = options.Temperature
		body.TopK = options.TopK
		body.TopP = options.TopP
		body.StopSequences = options.StopSequences

		payloadBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal payload body, %v", err)
		}

		return payloadBody, nil
	}

	// if config says ai21, use AI21PayloadBody
	if modelTLD == "ai21" {

		var body AI21PayloadBody
		body.Prompt = prompt
		body.Temperature = options.Temperature
		body.TopP = options.TopP
		body.MaxTokensToSample = options.MaxTokensToSample
		body.StopSequences = options.StopSequences

		payloadBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal payload body, %v", err)
		}

		return payloadBody, nil
	}

	// if config says cohere, use CoherePayloadBody
	if modelTLD == "cohere" {

		var body CoherePayloadBody
		body.Prompt = prompt
		body.Temperature = options.Temperature
		body.TopP = options.TopP
		body.TopK = float64(options.TopK)
		body.MaxTokensToSample = options.MaxTokensToSample
		body.StopSequences = options.StopSequences
		body.ReturnLiklihoods = options.ReturnLiklihoods
		body.Stream = options.Stream

		payloadBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal payload body, %v", err)
		}

		return payloadBody, nil
	}

	// if config says meta, use MetaPayloadBody
	if modelTLD == "meta" {

		var body MetaPayloadBody
		body.Prompt = prompt
		body.Temperature = options.Temperature
		body.TopP = options.TopP
		body.MaxTokensToSample = options.MaxTokensToSample

		payloadBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal payload body, %v", err)
		}

		return payloadBody, nil
	}

	// if config says stability, use StabilityPayloadBody
	if modelTLD == "stability" {

		var text StabilityTextPrompts
		text.Text = prompt

		var body StabilityPayloadBody
		body.Prompt = []StabilityTextPrompts{{Text: prompt}}
		body.Scale = 10
		body.Seed = 0
		body.Steps = 50

		payloadBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal payload body, %v", err)
		}

		return payloadBody, nil
	}

	return nil, fmt.Errorf("invalid model, %v", options.ModelID)

}

// SendToBedrockWithResponseStream is a function that sends a post request to Bedrock and returns the streaming response
func SendToBedrockWithResponseStream(prompt string, options Options) (*bedrockruntime.InvokeModelWithResponseStreamOutput, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(options.Region))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config, %v", err)
	}

	svc := bedrockruntime.NewFromConfig(cfg)

	accept := "*/*"
	modelId := options.ModelID
	contentType := "application/json"

	payloadBody, err := SerializePayload(prompt, options)
	if err != nil {
		return nil, fmt.Errorf("unable to serialize payload body, %v", err)
	}

	resp, err := svc.InvokeModelWithResponseStream(context.TODO(), &bedrockruntime.InvokeModelWithResponseStreamInput{
		Accept:      &accept,
		ModelId:     &modelId,
		ContentType: &contentType,
		Body:        []byte(string(payloadBody)),
	})
	if err != nil {
		return nil, fmt.Errorf("error from Bedrock, %v", err)
	}

	return resp, nil
}

// SendToBedrock is a function that sends a post request to Bedrock and returns the response
func SendToBedrock(prompt string, options Options) (*bedrockruntime.InvokeModelOutput, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(options.Region))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config, %v", err)
	}

	svc := bedrockruntime.NewFromConfig(cfg)

	accept := "*/*"
	modelId := options.ModelID
	contentType := "application/json"

	payloadBody, err := SerializePayload(prompt, options)
	if err != nil {
		return nil, fmt.Errorf("unable to serialize payload body, %v", err)
	}

	resp, err := svc.InvokeModel(context.TODO(), &bedrockruntime.InvokeModelInput{
		Accept:      &accept,
		ModelId:     &modelId,
		ContentType: &contentType,
		Body:        []byte(string(payloadBody)),
	})
	if err != nil {
		return nil, fmt.Errorf("error from Bedrock, %v", err)
	}

	return resp, nil
}
