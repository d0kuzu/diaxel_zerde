package main

import (
	"context"
	"diaxel/internal/api"
	appModule "diaxel/internal/app"
	"diaxel/internal/cleanup"
	"diaxel/internal/config"
	"diaxel/internal/grpc/db"
	"diaxel/internal/modules/llm"
	"diaxel/internal/modules/telegram"
	"diaxel/internal/modules/twilio"
	"log"
)

func main() {
	settings, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	grpcClient, err := db.New(settings.GRPCAddress)
	if err != nil {
		return
	}

	llmClient := llm.InitClient(settings.OpenaiApiKey, grpcClient)

	twilioClient := twilio.InitClient(settings.TwilioAccountSID, settings.TwilioAuthToken)

	tgOrchestrator := telegram.NewOrchestrator(llmClient, grpcClient, 5, 1000)
	go tgOrchestrator.Start(context.Background())

	cm := &cleanup.CleanupManager{}
	go cm.Start()

	app := appModule.NewApp(llmClient, twilioClient, grpcClient, settings, tgOrchestrator)

	api.RouterStart(app)
}
