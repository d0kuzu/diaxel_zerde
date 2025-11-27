package rest

import (
	"diaxel/api/infrastructure/controllers/ai_controllers"
	appModule "diaxel/app"
	"github.com/gin-gonic/gin"
)

func ChatRoutes(router *gin.Engine, app *appModule.App) {
	aiHandler := ai_controllers.NewAIHandler(app.Cfg, app.LLM)
	productGroup := router.Group("chat")
	{
		productGroup.POST("/send_message", aiHandler.SendMessage)
	}
}
