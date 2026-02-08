package ws

import (
	twilio "AISale/services/twillio"
	"diaxel/internal/config"
	"encoding/json"
	"log"
	"time"
)

func (c *Client) PollTwilio(chatID, accountSID, authToken string) {
	var lastMessageSID string

	for {
		messages, err := fetchMessagesFromTwilio(chatID, lastMessageSID, accountSID, authToken)
		if err != nil {
			log.Println("Twilio fetch error:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for _, m := range messages {
			var author string
			if m.From != config.BotNumber {
				author = "bot"
			} else {
				author = "client"
			}

			msg := Message{
				Author: author,
				Body:   m.Body,
			}

			data, err := json.Marshal(msg)
			if err != nil {
				log.Println("ws message json marshal error:", err)
			}

			c.Broadcast(chatID, data)

			lastMessageSID = m.Sid
		}
		time.Sleep(500 * time.Millisecond)
	}
}
