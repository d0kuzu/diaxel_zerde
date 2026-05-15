package llm

import (
	"diaxel/internal/grpc/db"
	"diaxel/internal/modules/calcom"

	"github.com/sashabaranov/go-openai"
)

type Client struct {
	client *openai.Client
	db     *db.Client
	calcom *calcom.Client
	model  string
}

func InitClient(openaiApiKey string, dbClient *db.Client, calcomClient *calcom.Client) *Client {
	return &Client{
		client: openai.NewClient(openaiApiKey),
		db:     dbClient,
		calcom: calcomClient,
		model:  "gpt-4o",
	}
}
