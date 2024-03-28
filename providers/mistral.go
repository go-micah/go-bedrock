package providers

type MistralMixtralInstructInvokeModelInput struct {
	Prompt            string   `json:"prompt"`
	Temperature       float64  `json:"temperature"`
	TopP              float64  `json:"top_p"`
	TopK              float64  `json:"top_k"`
	MaxTokensToSample int      `json:"max_tokens"`
	StopSequences     []string `json:"stop"`
}

type MistralMixtralInstructInvokeModelOutput struct {
	Text       string `json:"text"`
	StopReason string `json:"stop_reason,omitempty"`
}

type MistralMixtralInstructInvokeModelOutputs struct {
	Output []MistralMixtralInstructInvokeModelOutput `json:"outputs"`
}
