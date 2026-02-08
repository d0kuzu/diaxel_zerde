package webhook

import (
	appModule "diaxel/internal/app"
	"github.com/gin-gonic/gin"
)

func WebhookRoutes(router *gin.Engine, app *appModule.App) {
	aiHandler := NewAIHandler(app.Cfg, app.LLM)
	productGroup := router.Group("webhooks")
	{
		productGroup.POST("/telegram", aiHandler.SendMessage)
	}
}
