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
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
		return nil, err
	}

	log.Println("Connected to PostgreSQL")

	// Создаём таблицы, если их нет
	err = db.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
		&models.Assistant{},
		&models.Chat{},
		&models.Message{},
		&models.Analytics{},
	)
	if err != nil {
		log.Fatal("failed to migrate tables:", err)
		return nil, err
	}

	log.Println("Tables migrated successfully")
	return db, nil
}
