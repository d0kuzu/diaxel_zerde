package ws

import (
	appModule "diaxel/internal/app"

	"github.com/gin-gonic/gin"
)

func WSRoutes(router *gin.Engine, app *appModule.App) {
	wsHandler := NewWSHandler(app.Cfg, app.Db)
	productGroup := router.Group("websocket")
	{
		productGroup.GET("/get_conversation", wsHandler.ChatPolling)
	}
}
