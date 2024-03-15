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

type AnthropicClaudeMessagesInvokeModelInput struct {
	System           string                   `json:"system,omitempty"`
	Messages         []AnthropicClaudeMessage `json:"messages"`
	MaxTokens        int                      `json:"max_tokens"`
	Temperature      float64                  `json:"temperature"`
	TopK             int                      `json:"top_k"`
	TopP             float64                  `json:"top_p"`
	StopSequences    []string                 `json:"stop_sequences,omitempty"`
	AnthropicVersion string                   `json:"anthropic_version"`
}

type AnthropicClaudeMessage struct {
	Role    string                   `json:"role"`
	Content []AnthropicClaudeContent `json:"content"`
}

type AnthropicClaudeContent struct {
	Type   string                  `json:"type"`
	Text   string                  `json:"text,omitempty"`
	Source []AnthropicClaudeSource `json:"source,omitempty"`
}

type AnthropicClaudeSource struct {
	Type      string `json:"type"`
	MediaType string `json:"media_type"`
	Data      string `json:"data"`
}

type AnthropicClaudeMessagesInvokeModelOutput struct {
	Type  string `json:"type,omitempty"`
	Index int    `json:"index,omitempty"`
	Delta struct {
		Type string `json:"type,omitempty"`
		Text string `json:"text,omitempty"`
	} `json:"delta,omitempty"`
	Content    []AnthropicClaudeContent `json:"content,omitempty"`
	StopReason string                   `json:"stop_reason,omitempty"`
}
