package webhook

import (
	"diaxel/internal/config"
	"diaxel/internal/grpc/db"
	"diaxel/internal/modules/llm"
	"diaxel/internal/modules/token"
	"fmt"
	"net/http"
	"strconv"

	"diaxel/internal/modules/telegram"

	"github.com/gin-gonic/gin"
)

type AIHandler struct {
	cfg          *config.Settings
	LLM          *llm.Client
	db           *db.Client
	orchestrator *telegram.Orchestrator
}

func NewAIHandler(cfg *config.Settings, llmClient *llm.Client, db *db.Client, tgOrch *telegram.Orchestrator) *AIHandler {
	return &AIHandler{cfg: cfg, LLM: llmClient, db: db, orchestrator: tgOrch}
}

type RegisterBotRequest struct {
	Name string `json:"name"`
}

func (h *AIHandler) RegisterTelegramBot(c *gin.Context) {
	var req RegisterBotRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	userId := c.GetHeader("X-User-Id")
	if userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized request from gateway"})
		return
	}

	secureToken, err := token.GenerateSecureToken(h.cfg.TokenPrefix, h.cfg.TokenLength)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	assistant, err := h.db.CreateAssistant(req.Name, secureToken, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create assistant in database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Assistant registered", "assistant_id": assistant.Id, "token": secureToken})
}

type TelegramUpdate struct {
	UpdateID int64 `json:"update_id"`
	Message  struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
		From struct {
			ID int64 `json:"id"`
		} `json:"from"`
	} `json:"message"`
}

func (h *AIHandler) HandleTelegramWebhook(c *gin.Context) {
	assistantId := c.Param("assistant_id")

	var update TelegramUpdate

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "ignored - bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})

	if update.Message.Text == "" {
		return
	}

	userIdString := strconv.FormatInt(update.Message.From.ID, 10)
	if userIdString == "0" {
		return
	}

	assistant, err := h.db.GetAssistant(assistantId)
	if err != nil {
		fmt.Printf("HandleTelegramWebhook: Failed to get assistant: %v\n", err)
		return
	}

	if assistant.TelegramBotToken == "" {
		fmt.Printf("HandleTelegramWebhook: Assistant %s does not have a telegram bot token configured\n", assistantId)
		return
	}

	task := telegram.TelegramTask{
		BotToken:    assistant.TelegramBotToken,
		ChatID:      update.Message.Chat.ID,
		UserID:      userIdString,
		AssistantID: assistantId,
		UserMessage: update.Message.Text,
		UpdateID:    update.UpdateID,
	}

	h.orchestrator.Enqueue(task)
}

func (h *AIHandler) SendMessage(c *gin.Context) {
	userId := c.GetHeader("X-User-Id")
	assistantId := c.GetHeader("X-Assistant-Id")
	userMessage := c.PostForm("user_message")

	if userId == "" || assistantId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized request from gateway"})
		return
	}

	response, err := h.LLM.Conversation(c, userId, assistantId, userMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"answer": response})
}
