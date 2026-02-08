package main

import (
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"diaxel_zerde/database-service/models"
	"diaxel_zerde/database-service/proto"
	"diaxel_zerde/database-service/repository"
	"diaxel_zerde/database-service/server"
)

func main() {
	// Database connection
	dbHost := getEnv("POSTGRES_HOST", "localhost")
	dbPort := getEnv("POSTGRES_PORT", "5432")
	dbUser := getEnv("POSTGRES_USER", "postgres")
	dbPassword := getEnv("POSTGRES_PASSWORD", "password")
	dbName := getEnv("POSTGRES_DB", "diaxel")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Test database connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	defer sqlDB.Close()

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Database connected successfully")

	// Run migrations
	if err := runMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	refreshTokenRepo := repository.NewRefreshTokenRepository(db)
	chatRepo := repository.NewChatRepository(db)
	messageRepo := repository.NewMessageRepository(db)
	assistantRepo := repository.NewAssistantRepository(db)

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	databaseServer := server.NewDatabaseServer(userRepo, refreshTokenRepo, chatRepo, messageRepo, assistantRepo)

	proto.RegisterDatabaseServiceServer(grpcServer, databaseServer)

	// Enable reflection for development
	reflection.Register(grpcServer)

	// Start gRPC server
	port := getEnv("GRPC_PORT", "50051")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Starting gRPC server on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func runMigrations(db *gorm.DB) error {
	log.Println("Running database migrations...")

	// Import models to ensure they're registered with GORM
	// This will automatically create all tables based on the models
	err := db.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
		&models.Analytics{},
		&models.Assistant{},
		&models.Chat{},
		&models.Message{},
	)

	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}
