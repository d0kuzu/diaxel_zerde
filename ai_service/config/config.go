package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Settings struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
	Ssl        string

	OpenaiApiKey string

	ApiKey    string
	BaseID    string
	TableName string

	TwilioAccountSID string
	TwilioAuthToken  string
}

func LoadConfig() (*Settings, error) {
	godotenv.Load(".env")

	return &Settings{
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbPort:     os.Getenv("DB_PORT"),
		Ssl:        os.Getenv("DB_SSL"),

		OpenaiApiKey: os.Getenv("OPENAI_API_KEY"),

		ApiKey:    os.Getenv("API_KEY"),
		BaseID:    os.Getenv("BASE_ID"),
		TableName: os.Getenv("TABLE_NAME"),

		TwilioAccountSID: os.Getenv("TWILIO_ACCOUNT_SID"),
		TwilioAuthToken:  os.Getenv("TWILIO_AUTH_TOKEN"),
	}, nil
}
