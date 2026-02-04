package main

import (
	"log"

	"auth-service/internal/app"
	"auth-service/internal/config"
)

func main() {
	cfg := config.MustLoad()

	application := app.New(cfg)

	log.Fatal(application.Run())
}
