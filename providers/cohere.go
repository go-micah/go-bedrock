package providers

type CohereCommandInvokeModelInput struct {
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

type CohereCommandInvokeModelOutput struct {
	Generations []CohereCommandGeneration `json:"generations"`
}

type CohereCommandGeneration struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type CohereEmbedInvokeModelInput struct {
}

type CohereEmbedInvokeModelOutput struct {
}
