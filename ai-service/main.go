package main

import (
	"context"
	"log"

	_ "time/tzdata"

	"diaxel/internal/api"
	appModule "diaxel/internal/app"
	"diaxel/internal/cleanup"
	"diaxel/internal/config"
	"diaxel/internal/grpc/db"
	"diaxel/internal/modules/calcom"
	"diaxel/internal/modules/llm"
	"diaxel/internal/modules/telegram"
	"diaxel/internal/modules/twilio"
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

	calcomClient := calcom.New(settings.CalApiKey, settings.CalEventTypeID)

	llmClient := llm.InitClient(settings.OpenaiApiKey, grpcClient, calcomClient)

	twilioClient := twilio.InitClient()

	tgOrchestrator := telegram.NewOrchestrator(llmClient, grpcClient, 5, 1000)
	go tgOrchestrator.Start(context.Background())

	cm := &cleanup.CleanupManager{}
	go cm.Start()

	// followupListener := followup.NewListener(grpcClient, twilioClient)
	// go followupListener.Start(context.Background())

	app := appModule.NewApp(llmClient, twilioClient, grpcClient, settings, tgOrchestrator)

	api.RouterStart(app)
}
