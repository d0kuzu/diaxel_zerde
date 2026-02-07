package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort string

	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration

	AccessSecret  string
	RefreshSecret string
}

func MustLoad() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
		HTTPPort: os.Getenv("HTTP_PORT"),

		AccessTokenTTL:  15 * time.Minute,
		RefreshTokenTTL: 30 * 24 * time.Hour,

		AccessSecret:  os.Getenv("ACCESS_SECRET"),
		RefreshSecret: os.Getenv("REFRESH_SECRET"),
	}, nil
}
