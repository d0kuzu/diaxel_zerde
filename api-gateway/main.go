package main

import (
	"api-gateway/config"
	"api-gateway/server"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	s := server.NewServer(cfg)
	fmt.Println([]byte(cfg.AccessSecret))
	s.Run()
}
