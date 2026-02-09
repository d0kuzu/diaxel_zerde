package app

import (
	"diaxel/internal/config"
	"diaxel/internal/modules/llm"
	"diaxel/internal/modules/twilio"

	"gorm.io/gorm"
)

type LLM interface {
}

type Twilio interface {
}

type App struct {
	LLM    *llm.Client
	Twilio *twilio.Client
	Db     *gorm.DB
	Cfg    *config.Settings
}

func NewApp(llmClient *llm.Client, twilioClient *twilio.Client, db *gorm.DB, cfg *config.Settings) *App {
	return &App{
		LLM:    llmClient,
		Twilio: twilioClient,
		Db:     db,
		Cfg:    cfg,
	}
}
