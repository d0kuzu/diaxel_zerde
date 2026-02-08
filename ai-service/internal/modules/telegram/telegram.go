package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"diaxel/config"
	"diaxel/services/llm"
	pb "diaxel/services/telegram/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	llm      *llm.Client
	cfg      *config.Settings
	http     *http.Client
	dbClient pb.DatabaseServiceClient
}

type TelegramUpdate struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID        int64  `json:"id"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			ID        int64  `json:"id"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int64  `json:"date"`
		Text string `json:"text"`
	} `json:"message"`
}

type SendMessageRequest struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

func NewClient(llmClient *llm.Client, cfg *config.Settings) *Client {
	// Connect to database service
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to database service: %v", err)
	}

	dbClient := pb.NewDatabaseServiceClient(conn)

	return &Client{
		llm: llmClient,
		cfg: cfg,
		http: &http.Client{
			Timeout: 30 * time.Second,
		},
		dbClient: dbClient,
	}
}

func (c *Client) HandleWebhook(ctx context.Context, update TelegramUpdate) error {
	log.Printf("Received Telegram update: %+v", update)

	if update.Message.Text == "" {
		log.Println("Empty message text, skipping")
		return nil
	}

	userID := fmt.Sprintf("telegram_%d", update.Message.From.ID)
	chatID := fmt.Sprintf("telegram_%d", update.Message.Chat.ID)
	platform := "telegram"

	err := c.saveMessageWithRetry(ctx, chatID, "user", update.Message.Text, platform)
	if err != nil {
		log.Printf("Failed to save user message: %v", err)
		return fmt.Errorf("failed to save user message: %w", err)
	}

	ginCtx, _ := gin.CreateTestContext(nil)
	response, err := c.llm.Conversation(ginCtx, userID, update.Message.Text)
	if err != nil {
		log.Printf("Failed to generate LLM response: %v", err)
		return fmt.Errorf("failed to generate LLM response: %w", err)
	}

	err = c.saveMessageWithRetry(ctx, chatID, "assistant", response, platform)
	if err != nil {
		log.Printf("Failed to save assistant message: %v", err)
		return fmt.Errorf("failed to save assistant message: %w", err)
	}

	err = c.sendTelegramMessage(ctx, fmt.Sprintf("%d", update.Message.Chat.ID), response)
	if err != nil {
		log.Printf("Failed to send Telegram message: %v", err)
		return fmt.Errorf("failed to send Telegram message: %w", err)
	}

	log.Printf("Successfully processed Telegram message from user %s", userID)
	return nil
}

func (c *Client) saveMessageWithRetry(ctx context.Context, chatUserID, role, content, platform string) error {
	maxRetries := 3
	retryDelay := 1 * time.Second

	for i := 0; i < maxRetries; i++ {
		err := c.saveMessage(ctx, chatUserID, role, content, platform)
		if err == nil {
			return nil
		}

		log.Printf("Failed to save message (attempt %d/%d): %v", i+1, maxRetries, err)

		if i < maxRetries-1 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(retryDelay):
				retryDelay *= 2
			}
		}
	}

	return fmt.Errorf("failed to save message after %d attempts", maxRetries)
}

func (c *Client) saveMessage(ctx context.Context, chatUserID, role, content, platform string) error {
	req := &pb.SaveMessageRequest{
		ChatUserId: chatUserID,
		Role:       role,
		Content:    content,
		Platform:   platform,
	}

	_, err := c.dbClient.SaveMessage(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to save message via gRPC: %w", err)
	}

	log.Printf("Saved %s message for chat %s on platform %s", role, chatUserID, platform)
	return nil
}

func (c *Client) sendTelegramMessage(ctx context.Context, chatID, text string) error {
	if c.cfg.TelegramBotToken == "" {
		return fmt.Errorf("telegram bot token is not configured")
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", c.cfg.TelegramBotToken)

	reqBody := SendMessageRequest{
		ChatID:    chatID,
		Text:      text,
		ParseMode: "HTML",
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram API returned status %d: %s", resp.StatusCode, string(body))
	}

	log.Printf("Successfully sent Telegram message to chat %s", chatID)
	return nil
}

func (c *Client) ValidateWebhookSecret(secret string) bool {
	return c.cfg.TelegramWebhookSecret != "" && c.cfg.TelegramWebhookSecret == secret
}
