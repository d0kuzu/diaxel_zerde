package main

import (
	"diaxel/internal/api"
	appModule "diaxel/internal/app"
	"diaxel/internal/cleanup"
	"diaxel/internal/config"
	"diaxel/internal/grpc/db"
	"diaxel/internal/modules/llm"
	"diaxel/internal/modules/twilio"
	"log"
)

func main() {
	settings, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	grpcClient, err := db.New("database-service:50051")
	if err != nil {
		return
	}

	llmClient := llm.InitClient(settings.OpenaiApiKey)

	twilioClient := twilio.InitClient(settings.TwilioAccountSID, settings.TwilioAuthToken)

	//database.Connect(settings)

	cm := &cleanup.CleanupManager{}
	//cm.Add(database.Disconnect)
	go cm.Start()

	app := appModule.NewApp(llmClient, twilioClient, grpcClient, settings)

	api.RouterStart(app)
}
