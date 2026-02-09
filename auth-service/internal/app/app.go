package app

import (
	"auth-service/grpc/db"
	"auth-service/internal/api"
	"auth-service/internal/config"
	"auth-service/internal/service"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
	port   string
}

func New(cfg *config.Config) *App {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	grpcClient, err := db.New("database-service:50051")
	if err != nil {
		return nil
	}

	authService := service.NewAuthService(
		grpcClient,
		cfg,
	)

	// http
	api.RegisterRoutes(r, authService)

	return &App{
		router: r,
		port:   cfg.HTTPPort,
	}
}

func (a *App) Run() error {
	return a.router.Run(":" + a.port)
}
