package ws

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	chat string
}

func NewWSClient(conn *websocket.Conn, chatID string) *Client {
	return &Client{conn: conn, chat: chatID}
}

var (
	clients   = make(map[*Client]bool)
	clientsMu sync.Mutex
)

func RegisterClient(c *Client) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	clients[c] = true
}

func UnregisterClient(c *Client) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	delete(clients, c)
	c.conn.Close()
}

func (c *Client) Broadcast(chatID string, msg []byte) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	err := c.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("ws message write error:", err)
		c.conn.Close()
		delete(clients, c)
	}
}

func (c *Client) Listen(accountSID, authToken string) {
	defer UnregisterClient(c)

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("client disconnected:", err)
			return
		}
		log.Printf("[Operator → Twilio] Chat %s: %s\n", c.chat, string(msg))

		twilioClient := twilio.NewClient(accountSID, authToken)
		_, err = twilioClient.SendMessage(config.BotNumber, c.chat, string(msg))
		if err != nil {
			log.Println("ws twillio message send error:", err)
		}
	}
}
