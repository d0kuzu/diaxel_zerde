package followup

import (
	"context"
	"log"
	"time"

	"diaxel/internal/constants"
	"diaxel/internal/grpc/db"
	"diaxel/internal/modules/twilio"
)

type Listener struct {
	dbClient     *db.Client
	twilioClient *twilio.Client
}

func NewListener(dbClient *db.Client, twilioClient *twilio.Client) *Listener {
	return &Listener{
		dbClient:     dbClient,
		twilioClient: twilioClient,
	}
}

func (l *Listener) Start(ctx context.Context) {
	ticker := time.NewTicker(6 * time.Minute)
	defer ticker.Stop()

	log.Println("[Followup Listener] Started.")

	for {
		select {
		case <-ctx.Done():
			log.Println("[Followup Listener] Stopped.")
			return
		case <-ticker.C:
			l.processFollowups(ctx)
		}
	}
}

func (l *Listener) processFollowups(ctx context.Context) {
	//loc, err := time.LoadLocation("America/Winnipeg")
	//if err != nil {
	//	log.Printf("[Followup Listener] Error loading timezone: %v\n", err)
	//	return
	//}
	//
	//now := time.Now().In(loc)
	//hour := now.Hour()
	//
	//if hour < 9 || hour >= 18 {
	//	log.Printf("[Followup Listener] Outside working hours (current hour: %d in America/Winnipeg). Skipping.\n", hour)
	//	return
	//}

	log.Printf("[Followup Listener] Working")

	// We want to find chats inactive for 24 hours (24 * 60 * 60 seconds)
	inactiveDurationSeconds := int64(5 * 60)
	chats, err := l.dbClient.GetChatsForFollowup(inactiveDurationSeconds)
	if err != nil {
		log.Printf("[Followup Listener] Error getting chats for followup: %v\n", err)
		return
	}

	for _, chat := range chats {
		log.Printf("[Followup Listener] customer: %s", chat.CustomerId)
		// Only process if platform is twilio or we have customer id (phone number)
		// Assuming customerId is the phone number
		if chat.CustomerId == "" {
			continue
		}

		twilioConfig, err := l.dbClient.GetTwilioConfig(chat.AssistantId)
		if err != nil {
			log.Printf("[Followup Listener] Error getting twilio config for assistant %s: %v\n", chat.AssistantId, err)
			continue
		}

		err = l.twilioClient.SendMessage(
			ctx,
			twilioConfig.AccountSid,
			twilioConfig.AuthToken,
			twilioConfig.TwilioNumber,
			chat.CustomerId,
			constants.FollowupText,
		)

		if err != nil {
			log.Printf("[Followup Listener] Error sending followup to %s: %v\n", chat.CustomerId, err)
			continue
		}

		// Update chat as ended so we don't follow up again
		_, err = l.dbClient.UpdateChatIsEnd(chat.Id, true)
		if err != nil {
			log.Printf("[Followup Listener] Error updating chat %s to isEnd=true: %v\n", chat.Id, err)
		} else {
			log.Printf("[Followup Listener] Successfully sent followup and ended chat %s\n", chat.Id)
		}

		log.Printf("[Followup Listener] Work ended")
	}
}
