package chat

import (
	appModule "diaxel/internal/app"
	"github.com/gin-gonic/gin"
)

func ChatRoutes(router *gin.Engine, app *appModule.App) {
	aiHandler := NewAIHandler(app.Cfg, app.LLM)
	productGroup := router.Group("chat")
	{
		productGroup.POST("/send_message", aiHandler.SendMessage)
	}
}
