package main

import (
	"diaxel/api"
	appModule "diaxel/app"
	"diaxel/cleanup"
	"diaxel/config"
	"diaxel/services/llm"
	"log"
)

func main() {
	settings, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	llmClient := llm.InitClient(settings.OpenaiApiKey)

	cm := &cleanup.CleanupManager{}
	//cm.Add(chromeClient.Close)
	go cm.Start()

	app := appModule.NewApp(llmClient, settings)

	api.RouterStart(app)
}
