package server

import (
	"api-gateway/config"
	"api-gateway/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Run() {
	r := gin.Default()
	routes.SetupRoutes(r, s.cfg)

	addr := fmt.Sprintf(":%s", s.cfg.GatewayPort)
	fmt.Printf("Gateway running on %s\n", addr)
	if err := r.Run(addr); err != nil {
		panic(err)
	}
}
