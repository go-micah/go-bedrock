package providers

type MetaLlamaInvokeModelInput struct {
	Prompt            string  `json:"prompt"`
	Temperature       float64 `json:"temperature"`
	TopP              float64 `json:"top_p"`
	MaxTokensToSample int     `json:"max_gen_len"`
}

type MetaLlamaInvokeModelOutput struct {
	Generation string `json:"generation"`
}
