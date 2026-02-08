package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tr1ki/diaxel_zerde_master/database-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitPostgres инициализирует подключение к PostgreSQL и мигрирует модели
func InitPostgres() (*gorm.DB, error) {
	// Читаем настройки из .env или стандартные
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = "postgres"
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		password = "postgres"
	}
	dbName := os.Getenv("POSTGRES_DB")
	if dbName == "" {
		dbName = "diaxel_db"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)

	// Настройки логирования GORM
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	// Автоматическая миграция моделей
	err = db.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
		&models.Assistant{},
		&models.Chat{},
		&models.Message{},
		&models.Analytics{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate models: %w", err)
	}

	log.Println("Postgres initialized and models migrated successfully")
	return db, nil
}

// InitPostgresWithRetry инициализирует подключение с retry логикой
func InitPostgresWithRetry() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	maxRetries := 10
	retryDelay := 3 * time.Second

	for i := 0; i < maxRetries; i++ {
		log.Printf("Attempting to connect to PostgreSQL (attempt %d/%d)...", i+1, maxRetries)

		db, err = InitPostgres()
		if err == nil {
			log.Println("Successfully connected to PostgreSQL!")
			return db, nil
		}

		log.Printf("Failed to connect to PostgreSQL: %v", err)

		if i < maxRetries-1 {
			log.Printf("Waiting %v before retrying...", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	return nil, fmt.Errorf("failed to connect to PostgreSQL after %d attempts: %w", maxRetries, err)
}
