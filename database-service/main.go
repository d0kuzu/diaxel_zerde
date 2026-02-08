package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"diaxel_zerde/database-service/proto"
	"diaxel_zerde/database-service/repository"
	"diaxel_zerde/database-service/server"
)

func main() {
	// Database connection
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "diaxel_zerde")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test database connection
	if err := db.Ping(); err != nil {
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

	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	databaseServer := server.NewDatabaseServer(userRepo, refreshTokenRepo, chatRepo, messageRepo)

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

func runMigrations(db *sqlx.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Check if users table exists
	var exists bool
	err := db.GetContext(ctx, &exists, `
		SELECT EXISTS (
			SELECT FROM information_schema.tables 
			WHERE table_schema = 'public' 
			AND table_name = 'users'
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to check if tables exist: %w", err)
	}

	if !exists {
		log.Println("Running database migrations...")

		// Read and execute migration file
		migrationSQL := `
		-- Users table already exists, skip creation
		-- Refresh tokens table already exists, skip creation  
		-- Analytics table already exists, skip creation
		-- Assistants table already exists, skip creation
		-- Chats table already exists, skip creation
		-- Messages table already exists, skip creation
		
		-- Create indexes if they don't exist
		CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
		CREATE INDEX IF NOT EXISTS idx_assistants_user_id ON assistants(user_id);
		CREATE INDEX IF NOT EXISTS idx_chats_assistant_id ON chats(assistant_id);
		CREATE INDEX IF NOT EXISTS idx_chats_customer_id ON chats(customer_id);
		CREATE INDEX IF NOT EXISTS idx_messages_chat_user_id ON messages(chat_user_id);
		CREATE INDEX IF NOT EXISTS idx_refresh_tokens_token_hash ON refresh_tokens(token_hash);
		CREATE INDEX IF NOT EXISTS idx_refresh_tokens_user_id ON refresh_tokens(user_id);
		CREATE INDEX IF NOT EXISTS idx_refresh_tokens_expires_at ON refresh_tokens(expires_at);
		CREATE INDEX IF NOT EXISTS idx_analytics_assistant_id ON analytics(assistant_id);
		`

		if _, err := db.ExecContext(ctx, migrationSQL); err != nil {
			return fmt.Errorf("failed to run migrations: %w", err)
		}

		log.Println("Database migrations completed successfully")
	} else {
		log.Println("Database tables already exist, skipping migrations")
	}

	return nil
}
