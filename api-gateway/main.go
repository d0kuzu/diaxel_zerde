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
		log.Fatalf(err.Error())
	}

	grpcClient, err := db.New("localhost:50051")

	s := server.NewServer(cfg, grpcClient)
	s.Run()
}
