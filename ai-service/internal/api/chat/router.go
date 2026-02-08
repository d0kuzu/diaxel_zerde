package chat

import (
	"diaxel/internal/api/webhook"
	appModule "diaxel/internal/app"
	"github.com/gin-gonic/gin"
)

func ChatRoutes(router *gin.Engine, app *appModule.App) {
	aiHandler := webhook.NewAIHandler(app.Cfg, app.LLM)
	productGroup := router.Group("chat")
	{
		productGroup.POST("/send_message", aiHandler.SendMessage)
	}
}
