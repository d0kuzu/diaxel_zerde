package server

import (
	"api-gateway/config"
	"api-gateway/grpc/db"
	"api-gateway/routes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg *config.Config
	db  *db.Client
}

func NewServer(cfg *config.Config, db *db.Client) *Server {
	return &Server{cfg: cfg, db: db}
}

func (s *Server) Run() {
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	routes.SetupRoutes(r, s.cfg, db)

	addr := fmt.Sprintf(":%s", s.cfg.GatewayPort)
	fmt.Printf("Gateway running on %s\n", addr)
	if err := r.Run(addr); err != nil {
		panic(err)
	}
}
