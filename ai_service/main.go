package main

import (
	"diaxel/api"
	appModule "diaxel/app"
	"diaxel/cleanup"
	"diaxel/config"
	"diaxel/database"
	"diaxel/services/llm"
	"diaxel/services/twilio"
	"log"
)

func main() {
	settings, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	llmClient := llm.InitClient(settings.OpenaiApiKey)

	twilioClient := twilio.InitClient(settings.TwilioAccountSID, settings.TwilioAuthToken)

	database.Connect(settings)

	cm := &cleanup.CleanupManager{}
	cm.Add(database.Disconnect)
	go cm.Start()

	app := appModule.NewApp(llmClient, twilioClient, settings)

	api.RouterStart(app)
}
