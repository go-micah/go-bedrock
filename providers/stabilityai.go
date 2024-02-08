package providers

type StabilityAIStableDiffusionInvokeModelInput struct {
	Prompt []StabilityAIStableDiffusionTextPrompt `json:"text_prompts"`
	Scale  float64                                `json:"cfg_scale"`
	Steps  int                                    `json:"steps"`
	Seed   int                                    `json:"seed"`
}

type StabilityAIStableDiffusionTextPrompt struct {
	Text string `json:"text"`
}

type StabilityAIStableDiffusionInvokeModelOutput struct {
	Result    string                               `json:"result"`
	Artifacts []StabilityAIStableDiffusionArtifact `json:"artifacts"`
}

type StabilityAIStableDiffusionArtifact struct {
	Base64       string `json:"base64"`
	FinishReason string `json:"finishReason"`
}
