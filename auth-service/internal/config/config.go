package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort    string
	GRPCAddress string

	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration

	AccessSecret  string
	RefreshSecret string
}

func MustLoad() (*Config, error) {
	godotenv.Load(".env")

	grpcAddress := os.Getenv("GRPC_ADDRESS")
	if grpcAddress == "" {
		grpcAddress = "localhost:50051"
	}

	return &Config{
		HTTPPort:    os.Getenv("HTTP_PORT"),
		GRPCAddress: grpcAddress,

		AccessTokenTTL:  15 * time.Minute,
		RefreshTokenTTL: 30 * 24 * time.Hour,

		AccessSecret:  os.Getenv("ACCESS_SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
	}, nil
}
