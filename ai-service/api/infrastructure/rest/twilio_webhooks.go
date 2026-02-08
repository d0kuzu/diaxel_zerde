package rest

import (
	"diaxel/api/infrastructure/controllers/twilio"
	appModule "diaxel/app"
	"github.com/gin-gonic/gin"
)

func TwilioWebhookRoutes(router *gin.Engine, app *appModule.App) {
	twilioWebhookHandler := twilio.NewTwilioWebhookHandler(app.Cfg, app.LLM, app.Twilio)
	productGroup := router.Group("twilio")
	{
		productGroup.POST("/webhook", twilioWebhookHandler.HandleWebhook)
	}
}
