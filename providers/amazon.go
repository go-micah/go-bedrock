package providers

type AmazonTitanTextInvokeModelInput struct {
	Prompt string                          `json:"inputText"`
	Config AmazonTitanTextGenerationConfig `json:"textGenerationConfig"`
}

type AmazonTitanTextGenerationConfig struct {
	Temperature       float64  `json:"temperature"`
	TopP              float64  `json:"topP"`
	MaxTokensToSample int      `json:"maxTokenCount"`
	StopSequences     []string `json:"stopSequences"`
}

type AmazonTitanTextInvokeModelOutput struct {
	InputTextTokenCount int                      `json:"inputTextTokenCount"`
	Results             []AmazonTitanTextResults `json:"results"`
}

type AmazonTitanTextResults struct {
	TokenCount       int    `json:"tokenCount"`
	OutputText       string `json:"outputText"`
	CompletionReason string `json:"completionReason"`
}
