package providers

type AI21LabsJurassicInvokeModelInput struct {
	Prompt            string   `json:"prompt"`
	Temperature       float64  `json:"temperature"`
	TopP              float64  `json:"topP"`
	MaxTokensToSample int      `json:"maxTokens"`
	StopSequences     []string `json:"stopSequences"`
}

type AI21LabsJurrasicInvokeModelOutput struct {
	Completions []AI21LabsJurassicCompletions `json:"completions"`
}

type AI21LabsJurassicCompletions struct {
	Data AI21LabsJurassicData `json:"data"`
}

type AI21LabsJurassicData struct {
	Text string `json:"text"`
}
