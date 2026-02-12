package llm

import (
	"diaxel/internal/grpc/db"

	"github.com/sashabaranov/go-openai"
)

type Client struct {
	client      *openai.Client
	db          *db.Client
	assistantID string
	model       string
}

func InitClient(openaiApiKey string, dbClient *db.Client) *Client {
	return &Client{
		client: openai.NewClient(openaiApiKey),
		db:     dbClient,
		model:  "gpt-4o",
	}
}
