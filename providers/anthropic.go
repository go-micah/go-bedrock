package providers

type AnthropicClaudeInvokeModelInput struct {
	Prompt            string   `json:"prompt"`
	MaxTokensToSample int      `json:"max_tokens_to_sample"`
	Temperature       float64  `json:"temperature"`
	TopK              int      `json:"top_k"`
	TopP              float64  `json:"top_p"`
	StopSequences     []string `json:"stop_sequences"`
	AnthropicVersion  string   `json:"anthropic_version"`
}

type AnthropicClaudeInvokeModelOutput struct {
	Completion string `json:"completion"`
	StopReason string `json:"stop_reason"`
}
