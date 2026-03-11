package main

import (
	"api-gateway/config"
	"api-gateway/grpc/db"
	"api-gateway/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	grpcClient, err := db.New(cfg.GRPCAddress)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}

	s := server.NewServer(cfg, grpcClient)
	s.Run()
}
