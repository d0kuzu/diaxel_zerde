package chat

import (
	"diaxel/internal/config"
	"diaxel/internal/modules/llm"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AIHandler struct {
	cfg *config.Settings
	LLM *llm.Client
}

func NewAIHandler(cfg *config.Settings, llmClient *llm.Client) *AIHandler {
	return &AIHandler{cfg: cfg, LLM: llmClient}
}

type WebhookRequest struct {
	From string `json:"From"`
	Body string `json:"Body"`
}

func (h *AIHandler) SendMessage(c *gin.Context) {
	userId := c.PostForm("user_id")
	userMessage := c.PostForm("user_message")

	response, err := h.LLM.Conversation(c, userId, userMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": response})
}
