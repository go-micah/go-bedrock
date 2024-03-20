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

type AmazonTitanImageInvokeModelInput struct {
	TaskType              string                                                `json:"taskType"`
	ImageGenerationConfig AmazonTitanImageInvokeModelInputImageGenerationConfig `json:"imageGenerationConfig"`
	TextToImageParams     AmazonTitanImageInvokeModelInputTextToImageParams     `json:"textToImageParams"`
}

type AmazonTitanImageInvokeModelInputTextToImageParams struct {
	Text         string `json:"text"`
	NegativeText string `json:"negativeText,omitempty"`
}

type AmazonTitanImageInvokeModelInputImageGenerationConfig struct {
	NumberOfImages int     `json:"numberOfImages,omitempty"`
	Height         int     `json:"height,omitempty"`
	Width          int     `json:"width,omitempty"`
	Scale          float64 `json:"cfgScale,omitempty"`
	Seed           int     `json:"seed,omitempty"`
}

type AmazonTitanImageInvokeModelOutput struct {
	Images []string `json:"images"`
	Error  string   `json:"error"`
}
