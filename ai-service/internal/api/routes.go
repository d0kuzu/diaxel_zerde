package api

import (
	"diaxel/internal/api/analytics"
	"diaxel/internal/api/twilio"
	"diaxel/internal/api/webhook"
	appModule "diaxel/internal/app"
	"diaxel/internal/database"
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
	//ws.WSRoutes(r, app)
	//chat.ChatRoutes(r, app)
	analytics.SetupRoutes(r, database.GetDB())

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Router start error", err)
	}
}
