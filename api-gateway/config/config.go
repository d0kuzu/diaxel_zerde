package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GatewayPort    string
	UserServiceURL string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
		GatewayPort:    os.Getenv("GATEWAY_PORT"),
		UserServiceURL: os.Getenv("USER_SERVICE_URL"),
	}, nil
}
