package webhook

import (
	"bytes"
	"diaxel/internal/config"
	"diaxel/internal/grpc/db"
	"diaxel/internal/modules/llm"
	"diaxel/internal/modules/token"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	assistantRepos "diaxel/internal/database/models/repos"

	"github.com/gin-gonic/gin"
)

type AIHandler struct {
	cfg *config.Settings
	LLM *llm.Client
	db  *db.Client
}

func NewAIHandler(cfg *config.Settings, llmClient *llm.Client, db *db.Client) *AIHandler {
	return &AIHandler{cfg: cfg, LLM: llmClient, db: db}
}

type RegisterBotRequest struct {
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

func (h *AIHandler) RegisterTelegramBot(c *gin.Context) {
	var req RegisterBotRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	secureToken, err := token.GenerateSecureToken(h.cfg.TokenPrefix, h.cfg.TokenLength)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	assistant, err := h.db.CreateAssistant(req.Name, secureToken, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create assistant in database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Assistant registered", "assistant_id": assistant.Id, "token": secureToken})
}

type TelegramMessageReq struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func (h *AIHandler) HandleTelegramWebhook(c *gin.Context) {
	assistantId := c.Param("assistant_id")

	var updatePayload map[string]interface{}

	if err := c.ShouldBindJSON(&updatePayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	message, hasMessage := updatePayload["message"].(map[string]interface{})
	if !hasMessage {
		c.JSON(http.StatusOK, gin.H{"status": "ignored"})
		return
	}

	textInterface, hasText := message["text"]
	if !hasText {
		fmt.Println("Received non-text message")
		c.JSON(http.StatusOK, gin.H{"status": "ignored"})
		return
	}
	userText := textInterface.(string)

	chat, ok := message["chat"].(map[string]interface{})
	var chatID int64
	if ok {
		if idFloat, ok := chat["id"].(float64); ok {
			chatID = int64(idFloat)
		}
	}

	var userIdString string
	if from, ok := message["from"].(map[string]interface{}); ok {
		if idFloat, ok := from["id"].(float64); ok {
			userIdString = strconv.Itoa(int(idFloat))
		}
	}

	if userIdString == "" {
		fmt.Println("Could not extract user_id")
		c.JSON(http.StatusOK, gin.H{"status": "no user id"})
		return
	}

	assistant, err := assistantRepos.GetAssistant(assistantId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get assistant"})
		return
	}

	aiResponse, err := h.LLM.Conversation(c, userIdString, userText)

	if err != nil {
		fmt.Printf("AI Error: %v\n", err)
		c.JSON(http.StatusOK, gin.H{"status": "ai error"})
		return
	}

	err = h.sendTelegramMessage(assistant.BotToken, chatID, aiResponse)
	if err != nil {
		fmt.Printf("Failed to send to Telegram: %v\n", err)
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *AIHandler) sendTelegramMessage(token string, chatID int64, text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	reqBody := TelegramMessageReq{
		ChatID: chatID,
		Text:   text,
	}

	jsonData, _ := json.Marshal(reqBody)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status: %d", resp.StatusCode)
	}
	return nil
}

func (h *AIHandler) SendMessage(c *gin.Context) {
	userId := c.PostForm("user_id")
	userMessage := c.PostForm("user_message")
	response, err := h.LLM.Conversation(c, userId, userMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"answer": response})
}
