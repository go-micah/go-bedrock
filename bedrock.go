// Package bedrock is a wrapper around the Amazon Bedrock API
package bedrock

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type ModelInvoker interface {
	InvokeModel() (*bedrockruntime.InvokeModelOutput, error)
	InvokeModelWithResponseStream() (*bedrockruntime.InvokeModelWithResponseStreamOutput, error)
	GetText(resp *bedrockruntime.InvokeModelOutput) (string, error)
	GetDecodedImage(resp *bedrockruntime.InvokeModelOutput) ([]byte, error)
}

type AI21LabsJurassic struct {
	Region            string
	ModelId           string
	PromptRequest     string
	Temperature       float64
	TopP              float64
	MaxTokensToSample int
	StopSequences     []string
	Completions       []AI21LabsJurassicCompletions `json:"completions"`
}

type AI21LabsJurassicCompletions struct {
	Data AI21LabsJurassicData `json:"data"`
}

type AI21LabsJurassicData struct {
	Text string `json:"text"`
}

type AnthropicClaude struct {
	Region            string
	ModelId           string
	Prompt            string
	MaxTokensToSample int
	Temperature       float64
	TopK              int
	TopP              float64
	StopSequences     []string
	Completion        string `json:"completion"`
}

type AmazonTitanText struct {
}

type AmazonTitanEmbeddings struct {
}

type CohereCommand struct {
	Region            string
	ModelId           string
	Prompt            string
	Temperature       float64
	TopP              float64
	TopK              float64
	MaxTokensToSample int
	StopSequences     []string
	ReturnLiklihoods  string
	Stream            bool
	NumGenerations    int
	Generations       []CohereCommandGeneration `json:"generations"`
}

type CohereCommandGeneration struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type CohereEmbed struct {
}

type MetaLlama struct {
	Region            string
	ModelId           string
	Prompt            string  `json:"prompt"`
	Temperature       float64 `json:"temperature"`
	TopP              float64 `json:"top_p"`
	MaxTokensToSample int     `json:"max_gen_len"`
	Generation        string  `json:"generation"`
}

type StabilityAISD struct {
	Region    string
	ModelId   string
	Prompt    []StabilityAISDTextPrompts
	Scale     float64
	Steps     int
	Seed      int
	Result    string                  `json:"result"`
	Artifacts []StabilityAISDArtifact `json:"artifacts"`
}

type StabilityAISDTextPrompts struct {
	Text string `json:"text"`
}

type StabilityAISDArtifact struct {
	Base64       string `json:"base64"`
	FinishReason string `json:"finishReason"`
}

func sendToBedrock(payload []byte, modelId string, region string) (*bedrockruntime.InvokeModelOutput, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config, %v", err)
	}

	svc := bedrockruntime.NewFromConfig(cfg)

	accept := "*/*"
	contentType := "application/json"

	resp, err := svc.InvokeModel(context.TODO(), &bedrockruntime.InvokeModelInput{
		Accept:      &accept,
		ModelId:     &modelId,
		ContentType: &contentType,
		Body:        []byte(string(payload)),
	})
	if err != nil {
		return nil, fmt.Errorf("error from Bedrock, %v", err)
	}

	return resp, nil
}

func sendToBedrockWithResponseStream(payload []byte, modelId string, region string) (*bedrockruntime.InvokeModelWithResponseStreamOutput, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config, %v", err)
	}

	svc := bedrockruntime.NewFromConfig(cfg)

	accept := "*/*"
	contentType := "application/json"

	resp, err := svc.InvokeModelWithResponseStream(context.TODO(), &bedrockruntime.InvokeModelWithResponseStreamInput{
		Accept:      &accept,
		ModelId:     &modelId,
		ContentType: &contentType,
		Body:        []byte(string(payload)),
	})
	if err != nil {
		return nil, fmt.Errorf("error from Bedrock, %v", err)
	}

	return resp, nil
}

func (ai AI21LabsJurassic) InvokeModel() (*bedrockruntime.InvokeModelOutput, error) {

	type Payload struct {
		Prompt            string   `json:"prompt"`
		Temperature       float64  `json:"temperature"`
		TopP              float64  `json:"topP"`
		MaxTokensToSample int      `json:"maxTokens"`
		StopSequences     []string `json:"stopSequences"`
	}

	payload := Payload{
		Prompt:            ai.PromptRequest,
		MaxTokensToSample: ai.MaxTokensToSample,
		Temperature:       ai.Temperature,
		TopP:              ai.TopP,
		StopSequences:     ai.StopSequences,
	}

	payloadBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal payload body, %v", err)
	}

	resp, err := sendToBedrock(payloadBody, ai.ModelId, ai.Region)
	if err != nil {
		return nil, fmt.Errorf("error from Bedrock, %v", err)
	}

	return resp, nil
}

func (ac AnthropicClaude) InvokeModel() (*bedrockruntime.InvokeModelOutput, error) {

	type Payload struct {
		Prompt            string   `json:"prompt"`
		MaxTokensToSample int      `json:"max_tokens_to_sample"`
		Temperature       float64  `json:"temperature"`
		TopK              int      `json:"top_k"`
		TopP              float64  `json:"top_p"`
		StopSequences     []string `json:"stop_sequences"`
	}

	payload := Payload{
		Prompt:            ac.Prompt,
		MaxTokensToSample: ac.MaxTokensToSample,
		Temperature:       ac.Temperature,
		TopK:              ac.TopK,
		TopP:              ac.TopP,
		StopSequences:     ac.StopSequences,
	}

	payloadBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal payload body, %v", err)
	}

	resp, err := sendToBedrock(payloadBody, ac.ModelId, ac.Region)
	if err != nil {
		return nil, fmt.Errorf("error from Bedrock, %v", err)
	}
	return resp, nil
}

func (cc CohereCommand) InvokeModel() (*bedrockruntime.InvokeModelOutput, error) {

	type Payload struct {
		Prompt            string   `json:"prompt"`
		Temperature       float64  `json:"temperature"`
		TopP              float64  `json:"p"`
		TopK              float64  `json:"k"`
		MaxTokensToSample int      `json:"max_tokens"`
		StopSequences     []string `json:"stop_sequences"`
		ReturnLiklihoods  string   `json:"return_likelihoods"`
		Stream            bool     `json:"stream"`
		NumGenerations    int      `json:"num_generations"`
	}

	payload := Payload{
		Prompt:            cc.Prompt,
		Temperature:       cc.Temperature,
		TopK:              cc.TopK,
		TopP:              cc.TopP,
		MaxTokensToSample: cc.MaxTokensToSample,
		StopSequences:     cc.StopSequences,
		ReturnLiklihoods:  cc.ReturnLiklihoods,
		Stream:            false,
		NumGenerations:    cc.NumGenerations,
	}

	payloadBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal payload body, %v", err)
	}

	resp, err := sendToBedrock(payloadBody, cc.ModelId, cc.Region)
	if err != nil {
		return nil, fmt.Errorf("error from Bedrock, %v", err)
	}

	return resp, nil
}

func (ml MetaLlama) InvokeModel() (*bedrockruntime.InvokeModelOutput, error) {

	type Payload struct {
		Prompt            string  `json:"prompt"`
		Temperature       float64 `json:"temperature"`
		TopP              float64 `json:"top_p"`
		MaxTokensToSample int     `json:"max_gen_len"`
	}

	payload := Payload{
		Prompt:            ml.Prompt,
		Temperature:       ml.Temperature,
		TopP:              ml.TopP,
		MaxTokensToSample: ml.MaxTokensToSample,
	}

	payloadBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal payload body, %v", err)
	}

	resp, err := sendToBedrock(payloadBody, ml.ModelId, ml.Region)
	if err != nil {
		return nil, fmt.Errorf("error from Bedrock, %v", err)
	}

	return resp, nil
}

func (sd StabilityAISD) InvokeModel() (*bedrockruntime.InvokeModelOutput, error) {

	type Payload struct {
		Prompt []StabilityAISDTextPrompts `json:"text_prompts"`
		Scale  float64                    `json:"cfg_scale"`
		Steps  int                        `json:"steps"`
		Seed   int                        `json:"seed"`
	}

	payload := Payload{
		Prompt: sd.Prompt,
		Scale:  sd.Scale,
		Steps:  sd.Steps,
		Seed:   sd.Seed,
	}

	payloadBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal payload body, %v", err)
	}

	resp, err := sendToBedrock(payloadBody, sd.ModelId, sd.Region)
	if err != nil {
		return nil, fmt.Errorf("error from Bedrock, %v", err)
	}

	return resp, nil
}

func (ai AI21LabsJurassic) InvokeModelWithResponseStream() (*bedrockruntime.InvokeModelWithResponseStreamOutput, error) {
	return nil, fmt.Errorf("this model does not support streaming")
}

func (sd StabilityAISD) InvokeModelWithResponseStream() (*bedrockruntime.InvokeModelWithResponseStreamOutput, error) {
	return nil, fmt.Errorf("this model does not support streaming")
}

func (ac AnthropicClaude) InvokeModelWithResponseStream() (*bedrockruntime.InvokeModelWithResponseStreamOutput, error) {

	type Payload struct {
		Prompt            string   `json:"prompt"`
		MaxTokensToSample int      `json:"max_tokens_to_sample"`
		Temperature       float64  `json:"temperature"`
		TopK              int      `json:"top_k"`
		TopP              float64  `json:"top_p"`
		StopSequences     []string `json:"stop_sequences"`
	}

	payload := Payload{
		Prompt:            ac.Prompt,
		MaxTokensToSample: ac.MaxTokensToSample,
		Temperature:       ac.Temperature,
		TopK:              ac.TopK,
		TopP:              ac.TopP,
		StopSequences:     ac.StopSequences,
	}

	payloadBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal payload body, %v", err)
	}

	resp, err := sendToBedrockWithResponseStream(payloadBody, ac.ModelId, ac.Region)
	if err != nil {
		return nil, fmt.Errorf("error from Bedrock, %v", err)
	}

	return resp, nil

}

func (cc CohereCommand) InvokeModelWithResponseStream() (*bedrockruntime.InvokeModelWithResponseStreamOutput, error) {

	type Payload struct {
		Prompt            string   `json:"prompt"`
		Temperature       float64  `json:"temperature"`
		TopP              float64  `json:"p"`
		TopK              float64  `json:"k"`
		MaxTokensToSample int      `json:"max_tokens"`
		StopSequences     []string `json:"stop_sequences"`
		ReturnLiklihoods  string   `json:"return_likelihoods"`
		Stream            bool     `json:"stream"`
		NumGenerations    int      `json:"num_generations"`
	}

	payload := Payload{
		Prompt:            cc.Prompt,
		Temperature:       cc.Temperature,
		TopK:              cc.TopK,
		TopP:              cc.TopP,
		MaxTokensToSample: cc.MaxTokensToSample,
		StopSequences:     cc.StopSequences,
		ReturnLiklihoods:  cc.ReturnLiklihoods,
		Stream:            true,
		NumGenerations:    cc.NumGenerations,
	}

	payloadBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal payload body, %v", err)
	}

	resp, err := sendToBedrockWithResponseStream(payloadBody, cc.ModelId, cc.Region)
	if err != nil {
		return nil, fmt.Errorf("error from Bedrock, %v", err)
	}

	return resp, nil
}

func (ml MetaLlama) InvokeModelWithResponseStream() (*bedrockruntime.InvokeModelWithResponseStreamOutput, error) {

	type Payload struct {
		Prompt            string  `json:"prompt"`
		Temperature       float64 `json:"temperature"`
		TopP              float64 `json:"top_p"`
		MaxTokensToSample int     `json:"max_gen_len"`
	}

	payload := Payload{
		Prompt:            ml.Prompt,
		Temperature:       ml.Temperature,
		TopP:              ml.TopP,
		MaxTokensToSample: ml.MaxTokensToSample,
	}

	payloadBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal payload body, %v", err)
	}

	resp, err := sendToBedrockWithResponseStream(payloadBody, ml.ModelId, ml.Region)
	if err != nil {
		return nil, fmt.Errorf("error from Bedrock, %v", err)
	}

	return resp, nil
}

func (ai AI21LabsJurassic) GetText(resp *bedrockruntime.InvokeModelOutput) (string, error) {

	var jurassic AI21LabsJurassic

	err := json.Unmarshal(resp.Body, &jurassic)

	if err != nil {
		return "", fmt.Errorf("failed to unmarshal json, %v", err)
	}

	return jurassic.Completions[0].Data.Text, nil
}

func (ac AnthropicClaude) GetText(resp *bedrockruntime.InvokeModelOutput) (string, error) {

	var claude AnthropicClaude

	err := json.Unmarshal(resp.Body, &claude)

	if err != nil {
		return "", fmt.Errorf("failed to unmarshal json, %v", err)
	}

	return claude.Completion, nil
}

func (cc CohereCommand) GetText(resp *bedrockruntime.InvokeModelOutput) (string, error) {

	var command CohereCommand

	err := json.Unmarshal(resp.Body, &command)

	if err != nil {
		return "", fmt.Errorf("failed to unmarshal json, %v", err)
	}

	return command.Generations[0].Text, nil
}

func (ml MetaLlama) GetText(resp *bedrockruntime.InvokeModelOutput) (string, error) {

	var llama MetaLlama

	err := json.Unmarshal(resp.Body, &llama)

	if err != nil {
		return "", fmt.Errorf("failed to unmarshal json, %v", err)
	}

	return llama.Generation, nil
}

func (sd StabilityAISD) GetText(resp *bedrockruntime.InvokeModelOutput) (string, error) {
	return "", fmt.Errorf("this model does not support GetText")
}

func (sd StabilityAISD) GetDecodedImage(resp *bedrockruntime.InvokeModelOutput) ([]byte, error) {

	var stability StabilityAISD
	err := json.Unmarshal(resp.Body, &stability)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json, %v", err)
	}

	decoded, err := stability.Artifacts[0].decodeImage()

	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 response %v", err)
	}

	return decoded, nil

}

// decodeImage is a function that decodes the image from the response
func (a *StabilityAISDArtifact) decodeImage() ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(a.Base64)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
