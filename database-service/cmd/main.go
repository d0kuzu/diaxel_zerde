package main

import (
	"log"

	"github.com/tr1ki/diaxel_zerde_master/database-service/internal/db"
)

func main() {
	log.Println("Starting Database Service...")

	// Инициализация PostgreSQL с retry логикой
	database, err := db.InitPostgresWithRetry()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	log.Println("Database Service started successfully!")
	log.Printf("Database connection: %v", database)

	// Сервис готов к работе
	select {}
}
