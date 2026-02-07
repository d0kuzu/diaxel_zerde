package telegram

import (
	"net/http"

	"diaxel/config"
	"diaxel/services/llm"
	"diaxel/services/telegram"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB, llmClient *llm.Client, cfg *config.Settings) {
	telegramClient := telegram.NewClient(db, llmClient, cfg)

	r.POST("/webhook/telegram/:secret", func(c *gin.Context) {
		secret := c.Param("secret")

		if !telegramClient.ValidateWebhookSecret(secret) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid webhook secret"})
			return
		}

		var update telegram.TelegramUpdate
		if err := c.ShouldBindJSON(&update); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		err := telegramClient.HandleWebhook(c.Request.Context(), update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process update"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
