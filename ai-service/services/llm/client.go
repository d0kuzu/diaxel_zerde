package llm

import "github.com/sashabaranov/go-openai"

type Client struct {
	client *openai.Client
	agent  string
	model  string
}

func InitClient(openaiApiKey string) *Client {
	return &Client{
		client: openai.NewClient(openaiApiKey),
		model:  "gpt-4o",
	}
}
