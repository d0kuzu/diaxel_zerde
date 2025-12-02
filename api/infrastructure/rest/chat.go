package rest

import (
	"diaxel/api/infrastructure/controllers/ai"
	appModule "diaxel/app"
	"github.com/gin-gonic/gin"
)

func ChatRoutes(router *gin.Engine, app *appModule.App) {
	aiHandler := ai.NewAIHandler(app.Cfg, app.LLM)
	productGroup := router.Group("chat")
	{
		productGroup.POST("/send_message", aiHandler.SendMessage)
	}
}
