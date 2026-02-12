package chat

import (
	"diaxel/internal/config"
	"diaxel/internal/grpc/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	cfg *config.Settings
	db  *db.Client
}

func NewChatHandler(cfg *config.Settings, db *db.Client) *ChatHandler {
	return &ChatHandler{cfg: cfg, db: db}
}

type ChatResponse struct {
	ChatID       string                  `json:"chat_id"`
	AssistantID  string                  `json:"assistant_id"`
	CustomerID   string                  `json:"customer_id"`
	Messages     []*dbpb.MessageResponse `json:"messages"`
	MessageCount int32                   `json:"message_count"`
}

func (h *ChatHandler) GetAllChats(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil || page < 1 {
		page = 1
	}

	assistantID := c.Query("assistant_id")
	if assistantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "assistant_id is required"})
		return
	}

	chatsPerPage := int32(10)

	chats, totalCount, err := h.db.GetChatPage(assistantID, int32(page), chatsPerPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch chats", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"answer":      chats,
		"total_count": totalCount,
	})
}

func (h *ChatHandler) GetPagination(c *gin.Context) {
	assistantID := c.Query("assistant_id")
	if assistantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "assistant_id is required"})
		return
	}

	// Assuming we can get total count from GetChatPage with limit 0
	_, totalCount, err := h.db.GetChatPage(assistantID, 0, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch pagination", "details": err.Error()})
		return
	}

	pages := (totalCount + 9) / 10
	c.JSON(http.StatusOK, gin.H{"answer": pages})
}

func (h *ChatHandler) GetChat(c *gin.Context) {
	chatID := c.Query("chat")
	if chatID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chat parameter is required (chatID)"})
		return
	}

	messages, err := h.db.GetAllChatMessages(chatID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch chat messages", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": gin.H{
		"chat_id":  chatID,
		"messages": messages,
		"count":    len(messages),
	}})
}

func (h *ChatHandler) SearchChat(c *gin.Context) {
	searchTerm := c.Query("chat")
	assistantID := c.Query("assistant_id")
	if assistantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "assistant_id is required"})
		return
	}

	if searchTerm == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chat parameter is required (search term)"})
		return
	}

	chats, totalCount, err := h.db.SearchChatsByUser(assistantID, searchTerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to search chats", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"answer":      chats,
		"total_count": totalCount,
	})
}
