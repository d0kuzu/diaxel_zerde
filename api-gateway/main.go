package main

import (
	"api-gateway/config"
	"api-gateway/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	s := server.NewServer(cfg)
	s.Run()
}
