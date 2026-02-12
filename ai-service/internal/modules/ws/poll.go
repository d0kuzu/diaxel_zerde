package ws

import (
	"encoding/json"
	"log"
	"time"
)

func (c *Client) PollLocalDB(chatID string) {
	var lastMessageCount int

	for {
		chat, err := chat_repos.CheckIfExist(chatID)
		if err != nil {
			log.Println("DB fetch error:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		currentMessageCount := len(chat.Messages)

		if currentMessageCount > lastMessageCount {
			for i := lastMessageCount; i < currentMessageCount; i++ {
				msg := chat.Messages[i]

				var author string
				if msg.Role == "user" {
					author = "client"
				} else {
					author = "bot"
				}

				wsMsg := Message{
					Author: author,
					Body:   msg.Content,
				}

				data, err := json.Marshal(wsMsg)
				if err != nil {
					log.Println("ws message json marshal error:", err)
					continue
				}

				c.Broadcast(chatID, data)
			}
			lastMessageCount = currentMessageCount
		}

		time.Sleep(500 * time.Millisecond)
	}
}
