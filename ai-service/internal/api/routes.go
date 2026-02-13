package api

import (
	"diaxel/internal/api/chat"
	"diaxel/internal/api/twilio"
	"diaxel/internal/api/webhook"
	"diaxel/internal/api/ws"
	appModule "diaxel/internal/app"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterStart(app *appModule.App) {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		MaxAge:       12 * 60 * 60,
	}))

	webhook.WebhookRoutes(r, app)
	twilio.TwilioWebhookRoutes(r, app)
	ws.WSRoutes(r, app)
	chat.ChatRoutes(r, app)

	// analyticsService := analytics.NewAnalyticsService(app.Db)
	// analyticsAPI.SetupRoutes(r, analyticsService)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Router start error", err)
	}
}
