package assistant

import (
	"diaxel/internal/config"
	"diaxel/internal/grpc/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssistantHandler struct {
	cfg *config.Settings
	db  *db.Client
}

func NewAssistantHandler(cfg *config.Settings, db *db.Client) *AssistantHandler {
	return &AssistantHandler{cfg: cfg, db: db}
}

func (h *AssistantHandler) GetAssistants(c *gin.Context) {
	userId := c.GetHeader("X-User-Id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-User-Id header is required"})
		return
	}

	assistants, err := h.db.GetAssistantsByUserID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch assistants", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"answer": assistants,
	})
}
