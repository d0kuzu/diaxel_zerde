package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/tr1ki/diaxel_zerde_master/database-service/internal/models"
)

func InitPostgres() (*gorm.DB, error) {
	// Берём настройки из переменных окружения
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	// Авто миграция всех моделей
	err = db.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
		&models.Assistant{},
		&models.Chat{},
		&models.Message{},
		&models.Analytics{},
	)
	if err != nil {
		return nil, fmt.Errorf("auto migrate failed: %w", err)
	}

	log.Println("Postgres connected and migrated")
	return db, nil
}
