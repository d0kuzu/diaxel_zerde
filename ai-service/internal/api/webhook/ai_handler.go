package webhook

import (
	"diaxel/internal/config"
	"diaxel/internal/grpc/db"
	"diaxel/internal/modules/llm"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AIHandler struct {
	cfg *config.Settings
	LLM *llm.Client
	db  *db.Client
}

func NewAIHandler(cfg *config.Settings, llmClient *llm.Client, db *db.Client) *AIHandler {
	return &AIHandler{cfg: cfg, LLM: llmClient, db: db}
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

func (h *AIHandler) Test(c *gin.Context) {
	h.db.GetStats()

	c.JSON(http.StatusOK, gin.H{"answer": "ok"})
}
