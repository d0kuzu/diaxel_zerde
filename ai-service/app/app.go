package app

import (
	"diaxel/config"
	"diaxel/services/llm"
	"diaxel/services/twilio"
)

type App struct {
	LLM    *llm.Client
	Twilio *twilio.Client
	Cfg    *config.Settings
}

func NewApp(llmClient *llm.Client, twilioClient *twilio.Client, cfg *config.Settings) *App {
	return &App{
		LLM:    llmClient,
		Twilio: twilioClient,
		Cfg:    cfg,
	}
}
