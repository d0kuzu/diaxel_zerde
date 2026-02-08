package webhook

import (
	appModule "diaxel/internal/app"
	"github.com/gin-gonic/gin"
)

func WebhookRoutes(router *gin.Engine, app *appModule.App) {
	aiHandler := NewAIHandler(app.Cfg, app.LLM, app.Db)
	productGroup := router.Group("webhooks")
	{
		productGroup.POST("/telegram", aiHandler.SendMessage)
		productGroup.POST("/test", aiHandler.Test)
	}
}
