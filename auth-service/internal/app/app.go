package app

import (
	"auth-service/internal/api"
	"auth-service/internal/config"
	"auth-service/internal/repository"
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

	// repositories
	userRepo := repository.NewUserRepo()
	refreshRepo := repository.NewRefreshRepo()

	// services
	authService := service.NewAuthService(
		userRepo,
		refreshRepo,
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
