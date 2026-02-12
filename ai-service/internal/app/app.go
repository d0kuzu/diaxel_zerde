package app

import (
	"diaxel/internal/config"
	"diaxel/internal/grpc/db"
	"diaxel/internal/modules/llm"
	"diaxel/internal/modules/twilio"
)

type LLM interface {
}

type Twilio interface {
}

type App struct {
	LLM    *llm.Client
	Twilio *twilio.Client
	Db     *db.Client
	Cfg    *config.Settings
}

func NewApp(llmClient *llm.Client, twilioClient *twilio.Client, db *db.Client, cfg *config.Settings) *App {
	return &App{
		LLM:    llmClient,
		Twilio: twilioClient,
		Db:     db,
		Cfg:    cfg,
	}
}
