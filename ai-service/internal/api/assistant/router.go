package assistant

import (
	appModule "diaxel/internal/app"

	"github.com/gin-gonic/gin"
)

func AssistantRoutes(router *gin.Engine, app *appModule.App) {
	h := NewAssistantHandler(app.Cfg, app.Db)

	assistantGroup := router.Group("assistants")
	{
		assistantGroup.GET("/list", h.GetAssistants)
	}
}
