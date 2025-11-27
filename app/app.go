package app

import (
	"diaxel/config"
	"diaxel/services/llm"
)

type App struct {
	LLM *llm.Client
	Cfg *config.Settings
}

func NewApp(llmClient *llm.Client, cfg *config.Settings) *App {
	return &App{
		LLM: llmClient,
		Cfg: cfg,
	}
}
