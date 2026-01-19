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
	r.Use(auth.AuthMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	users := r.Group("/users")
	users.Any("/*any", proxy.NewReverseProxy(cfg.UserServiceURL, "/users"))

	// Здесь можно добавить аггрегацию запросов позже, например:
	// r.GET("/dashboard", aggregateHandler)
}
