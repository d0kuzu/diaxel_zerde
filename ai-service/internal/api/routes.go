package api

import (
	"diaxel/internal/api/analytics"
	"diaxel/internal/api/chat"
	"diaxel/internal/api/telegram"
	"diaxel/internal/api/twilio"
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

	chat.ChatRoutes(r, app)
	twilio.TwilioWebhookRoutes(r, app)

	telegram.SetupRoutes(r, database.GetDB(), app.LLM, app.Cfg)
	analytics.SetupRoutes(r, database.GetDB())

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Router start error", err)
	}
}
