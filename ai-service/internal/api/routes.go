package api

import (
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

	// TODO: Fix webhook and twilio routes
	// webhook.WebhookRoutes(r, app)
	// twilio.TwilioWebhookRoutes(r, app)
	// ws.WSRoutes(r, app)
	// chat.ChatRoutes(r, app)

	// TODO: Initialize analytics service with gRPC client
	// analyticsService := analytics.NewAnalyticsService(app.Db)
	// TODO: Fix analytics routes
	// analyticsAPI.SetupRoutes(r, analyticsService)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Router start error", err)
	}
}
