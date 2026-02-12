package ws

import (
	"diaxel/internal/config"
	"diaxel/internal/grpc/db"
	"diaxel/internal/modules/ws"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WSHandler struct {
	cfg *config.Settings
	db  *db.Client
}

func NewWSHandler(cfg *config.Settings, db *db.Client) *WSHandler {
	return &WSHandler{cfg: cfg, db: db}
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

	client := ws.NewWSClient(conn, chatID, h.db)
	ws.RegisterClient(client)
	log.Println("new client connected to chat", chatID)

	go client.PollLocalDB(chatID)

	go client.Listen()
}
