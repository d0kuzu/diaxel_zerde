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
	}

	private := r.Group("/")
	private.Use(auth.AuthMiddleware())
	{
		private.Any("/users/*any",
			proxy.NewReverseProxy(cfg.UserServiceURL, "/users"),
		)
	}
}
