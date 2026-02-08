package routes

import (
	"api-gateway/config"
	"api-gateway/middleware/auth"
	"api-gateway/middleware/logger"
	"api-gateway/proxy"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config) {
	r.Use(gin.LoggerWithFormatter(logger.Formatter))

	public := r.Group("/")
	{
		public.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})

		public.Any("/auth/*any",
			proxy.NewReverseProxy(cfg.AuthServiceURL, "/auth"),
		)

		public.Any("/webhooks/telegram/*any",
			proxy.NewReverseProxy(cfg.AIServiceURL, "/webhooks/telegram"),
		)
	}

	userPrivate := r.Group("/")
	userPrivate.Use(auth.UserMiddleware([]byte(cfg.AccessSecret)))
	{
		userPrivate.Any("/users/*any",
			proxy.NewReverseProxy(cfg.UserServiceURL, "/users"),
		)

		userPrivate.Any("/api/analytics/*any",
			proxy.NewReverseProxy(cfg.AIServiceURL, "/api/analytics"),
		)
	}

	servicePrivate := r.Group("/")
	servicePrivate.Use(auth.ServiceMiddleware([]byte(cfg.TelegramServiceSecret), "telegram-service", "ai-service"))
	{
		servicePrivate.Any("/internal/analytics/*any",
			proxy.NewReverseProxy(cfg.AIServiceURL, "/api/analytics"),
		)
	}
}
