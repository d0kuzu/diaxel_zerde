package llm

import (
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

func (c *Client) GetAnswer(ctx *gin.Context, messages []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {
	response, err := c.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    c.model,
		Messages: messages,
	})
	if err != nil {
		return openai.ChatCompletionResponse{}, err
	}

	return response, nil
}
