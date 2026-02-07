package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GatewayPort string

	UserServiceURL string
	AuthServiceURL string
	AIServiceURL   string

	TelegramWebhook string

	AccessSecret          string
	TelegramServiceSecret string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
		GatewayPort: os.Getenv("GATEWAY_PORT"),

		UserServiceURL: os.Getenv("USER_SERVICE_URL"),
		AuthServiceURL: os.Getenv("AUTH_SERVICE_URL"),
		AIServiceURL:   os.Getenv("AI_SERVICE_URL"),

		TelegramWebhook: os.Getenv("TELEGRAM_WEBHOOK"),

		AccessSecret:          os.Getenv("ACCESS_SECRET"),
		TelegramServiceSecret: os.Getenv("TELEGRAM_SERVICE_SECRET"),
	}, nil
}
