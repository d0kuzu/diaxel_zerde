package llm

import (
	"diaxel/internal/grpc/db"
	"diaxel/internal/modules/calcom"
	"diaxel/internal/modules/campuslogin"

	"github.com/sashabaranov/go-openai"
)

type Client struct {
	client      *openai.Client
	db          *db.Client
	calcom      *calcom.Client
	campuslogin *campuslogin.Client
	model       string
}

func InitClient(openaiApiKey string, dbClient *db.Client, calcomClient *calcom.Client, campusloginClient *campuslogin.Client) *Client {
	return &Client{
		client:      openai.NewClient(openaiApiKey),
		db:          dbClient,
		calcom:      calcomClient,
		campuslogin: campusloginClient,
		model:       "gpt-4o",
	}
}
