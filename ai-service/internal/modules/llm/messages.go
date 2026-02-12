package llm

import (
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
)

// Message represents a chat message (moved from database models)
type Message struct {
	ChatID  uuid.UUID `json:"chat_id"`
	Role    string    `json:"role"`
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}

// GetHistory retrieves all messages for the user.
func (c *Client) GetHistory(customerId string) ([]Message, error) {
	// 1. Check if chat exists
	chatResp, err := c.db.GetLatestChatByCustomer(c.assistantID, customerId)
	if err != nil {
		return nil, err
	}

	// If chat ID is empty (not found), return empty history
	if chatResp.Id == "" {
		return []Message{}, nil
	}

	// 2. Fetch all messages
	messagesResp, err := c.db.GetAllChatMessages(chatResp.Id)
	if err != nil {
		return nil, err
	}

	var history []Message
	for _, msg := range messagesResp {
		history = append(history, Message{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	return history, nil
}

func (c *Client) GetAllChats() ([]string, error) {
	// TODO: Implement using gRPC calls to database service if needed
	return []string{}, nil
}

func (c *Client) GetMessages(customerId string) ([]openai.ChatCompletionMessage, error) {
	history, err := c.GetHistory(customerId)
	if err != nil {
		return nil, err
	}

	messages, err := c.ConvertToOpenaiMessage(history)
	if err != nil {
		return nil, err
	}

	// TODO DOKUZU: этот метод должен добавлять системные сообщение как последнее сообщение
	// Fetch system prompt from DB via assistant configuration
	startMessages := c.StartMessages()
	messages = append(messages, startMessages...)

	return messages, nil
}

func (c *Client) StartMessages() []openai.ChatCompletionMessage {
	log.Printf("Принял системный промпт")
	// TODO DOKUZU: system messages нужно брать из бд, у ассистента должно быть поле для этого

	systemPrompt := "You are a helpful assistant."

	// Try to get dynamic prompt from assistant config via gRPC
	assistant, err := c.db.GetAssistant(c.assistantID)
	if err == nil && assistant.Configuration != "" {
		systemPrompt = assistant.Configuration
	}

	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
	}
}

func (c *Client) AddMessage(messages *[]openai.ChatCompletionMessage, role string, message string) {
	*messages = append(*messages, openai.ChatCompletionMessage{Role: role, Content: message})
}

func (c *Client) SaveMessages(customerId string, messages []openai.ChatCompletionMessage) error {
	log.Printf("Saving messages for user %s", customerId)

	// 1. Get or create chat
	chatResp, err := c.db.GetLatestChatByCustomer(c.assistantID, customerId)
	if err != nil {
		return err
	}

	chatID := chatResp.Id
	if chatID == "" {
		newChat, err := c.db.CreateChat(c.assistantID, customerId, "openai")
		if err != nil {
			return err
		}
		chatID = newChat.Id
	}

	// 2. Filter and save new messages
	for _, msg := range messages {
		// TODO DOKUZU: при сохранении, system сообщения не добавляются
		if msg.Role == openai.ChatMessageRoleSystem {
			continue
		}

		_, err := c.db.SaveMessage(chatID, msg.Role, msg.Content, "openai")
		if err != nil {
			log.Printf("Failed to save message: %v", err)
		}
	}

	return nil
}

func (c *Client) ConvertToMessage(customerId string, messages []openai.ChatCompletionMessage) []Message {
	var messagesArray []Message

	for _, message := range messages {
		messagesArray = append(messagesArray, Message{
			Role:    message.Role,
			Content: message.Content,
			Time:    time.Now(),
		})
	}

	return messagesArray
}

func (c *Client) ConvertToOpenaiMessage(arrayMessages []Message) ([]openai.ChatCompletionMessage, error) {
	var messages []openai.ChatCompletionMessage

	for _, message := range arrayMessages {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    message.Role,
			Content: message.Content,
		})
	}

	return messages, nil
}

func (c *Client) RemoveSystemMessages(messages *[]openai.ChatCompletionMessage) {
	var otherMessages []openai.ChatCompletionMessage

	for _, message := range *messages {
		if message.Role != "system" || strings.Contains(message.Content, "#function_call") {
			otherMessages = append(otherMessages, message)
		}
	}

	*messages = otherMessages
}
