package chat

import (
	"diaxel/internal/config"
	"diaxel/internal/database/models"
	"diaxel/internal/database/models/repos/chat_repos"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ChatHandler struct {
	cfg *config.Settings
	db  *gorm.DB
}

func NewChatHandler(cfg *config.Settings, db *gorm.DB) *ChatHandler {
	return &ChatHandler{cfg: cfg, db: db}
}

type ChatResponse struct {
	UserID       string           `json:"user_id"`
	Messages     []models.Message `json:"messages"`
	IsClient     bool             `json:"is_client"`
	MessageCount int              `json:"message_count"`
}

func (h *ChatHandler) GetAllChats(c *gin.Context) {
	page, err := strconv.ParseInt(c.Query("page"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid page parameter"})
		return
	}

	offset := (page - 1) * 10
	var chats []models.Chat

	if err := h.db.Preload("Messages").Offset(int(offset)).Limit(10).Find(&chats).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch chats", "details": err.Error()})
		return
	}

	var response []ChatResponse
	for _, chat := range chats {
		response = append(response, ChatResponse{
			UserID:       chat.UserID.String(),
			Messages:     chat.Messages,
			IsClient:     chat.IsClient,
			MessageCount: len(chat.Messages),
		})
	}

	c.JSON(200, gin.H{"answer": response})
}

func (h *ChatHandler) GetPagination(c *gin.Context) {
	var count int64

	if err := h.db.Model(&models.Chat{}).Count(&count).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	pages := (count + 9) / 10 // ceiling division
	c.JSON(200, gin.H{"answer": pages})
}

func (h *ChatHandler) GetChat(c *gin.Context) {
	userID := c.Query("chat")
	if userID == "" {
		c.JSON(400, gin.H{"error": "chat parameter is required"})
		return
	}

	chat, err := chat_repos.CheckIfExist(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch chat", "details": err.Error()})
		return
	}

	response := ChatResponse{
		UserID:       chat.UserID.String(),
		Messages:     chat.Messages,
		IsClient:     chat.IsClient,
		MessageCount: len(chat.Messages),
	}

	c.JSON(200, gin.H{"answer": response})
}

func (h *ChatHandler) SearchChat(c *gin.Context) {
	searchTerm := c.Query("chat")
	if searchTerm == "" {
		c.JSON(400, gin.H{"error": "chat parameter is required"})
		return
	}

	var chats []models.Chat

	if err := h.db.Preload("Messages").Where("user_id ILIKE ?", "%"+searchTerm+"%").Find(&chats).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to search chats", "details": err.Error()})
		return
	}

	var response []ChatResponse
	for _, chat := range chats {
		response = append(response, ChatResponse{
			UserID:       chat.UserID.String(),
			Messages:     chat.Messages,
			IsClient:     chat.IsClient,
			MessageCount: len(chat.Messages),
		})
	}

	c.JSON(200, gin.H{"answer": response})
}
