[![Go Reference](https://pkg.go.dev/badge/github.com/go-micah/go-bedrock.svg)](https://pkg.go.dev/github.com/go-micah/go-bedrock)

# go-bedrock

A wrapper around the Amazon Bedrock API written in Go

## Use

```go
prompt := "Please write me a short poem about a chicken"

// prepare payload for Anthropic Claude v2
claude := bedrock.AnthropicClaude{
    Region:            "us-east-1",
    ModelId:           "anthropic.claude-v2",
    Prompt:            "Human: \n\nHuman: " + prompt + "\n\nAssistant:",
    MaxTokensToSample: 500,
    TopP:              0.999,
    TopK:              250,
    Temperature:       1,
    StopSequences:     []string{`"\n\nHuman:\"`},
}

fmt.Println("Sending prompt to Anthropic Claude v2")

resp, err := claude.InvokeModel()
if err != nil {
    log.Fatal("error", err)
}

text, err := claude.GetText(resp)
if err != nil {
    log.Fatal("error", err)
}
fmt.Println(text)
```
