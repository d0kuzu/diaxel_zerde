package main

import (
	"log"

	"auth-service/internal/app"
	"auth-service/internal/config"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	application := app.New(cfg)

	log.Fatal(application.Run())
}
