package ws

import (
	"diaxel/internal/config"
	"diaxel/internal/modules/ws"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WSHandler struct {
	cfg *config.Settings
}

func NewWSHandler(cfg *config.Settings) *WSHandler {
	return &WSHandler{cfg: cfg}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (h *WSHandler) ChatPolling(c *gin.Context) {
	chatID := c.Query("chat")
	if chatID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chat id required"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}

	client := ws.NewWSClient(conn, chatID)
	ws.RegisterClient(client)
	log.Println("new client connected to chat", chatID)

	go client.PollTwilio(chatID, h.cfg.AccountSID, h.cfg.AuthToken)

	go client.Listen(h.cfg.AccountSID, h.cfg.AuthToken)
}
