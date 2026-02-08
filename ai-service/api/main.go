package api

import (
	"diaxel/api/infrastructure/rest"
	appModule "diaxel/app"
	"diaxel/database"
	"diaxel/internal/analytics"
	"diaxel/services/webhooks/telegram"
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

	rest.ChatRoutes(r, app)
	rest.TwilioWebhookRoutes(r, app)

	telegram.SetupRoutes(r, database.GetDB(), app.LLM, app.Cfg)
	analytics.SetupRoutes(r, database.GetDB())

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Router start error", err)
	}
}
