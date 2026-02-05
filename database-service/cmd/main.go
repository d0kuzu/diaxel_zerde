package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("database-service started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("database-service shutting down")
}
