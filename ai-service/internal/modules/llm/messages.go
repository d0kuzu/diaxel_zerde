package llm

import (
	"diaxel/internal/constants"
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

// Mock functions for now - will be replaced with gRPC calls
func GetHistory(userId string) ([]Message, error) {
	// TODO: Implement using gRPC calls to database service
	return []Message{}, nil
}

func GetAllChats() ([]string, error) {
	// TODO: Implement using gRPC calls to database service
	return []string{}, nil
}

func GetMessages(userId string) ([]openai.ChatCompletionMessage, error) {
	// TODO: Implement using gRPC calls to database service
	var messages []openai.ChatCompletionMessage

	startMessages := StartMessages()
	messages = append(messages, startMessages...)

	return messages, nil
}

func StartMessages() []openai.ChatCompletionMessage {
	log.Printf("Принял системный промпт")
	return constants.SystemMessages
}

func AddMessage(messages *[]openai.ChatCompletionMessage, role string, message string) {
	*messages = append(*messages, openai.ChatCompletionMessage{Role: role, Content: message})
}

func SaveMessages(userId string, messages []openai.ChatCompletionMessage) error {
	// TODO: Implement using gRPC calls to database service
	log.Printf("Saving messages for user %s", userId)
	return nil
}

func ConvertToMessage(userId string, messages []openai.ChatCompletionMessage) []Message {
	var messagesArray []Message

	// Convert userId to UUID
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		return messagesArray
	}

	for _, message := range messages {
		messagesArray = append(messagesArray, Message{
			ChatID:  userUUID,
			Role:    message.Role,
			Content: message.Content,
			Time:    time.Now(),
		})
	}

	return messagesArray
}

func ConvertToOpenaiMessage(arrayMessages []Message) ([]openai.ChatCompletionMessage, error) {
	var messages []openai.ChatCompletionMessage

	for _, message := range arrayMessages {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    message.Role,
			Content: message.Content,
		})
	}

	return messages, nil
}

func RemoveSystemMessages(messages *[]openai.ChatCompletionMessage) {
	var otherMessages []openai.ChatCompletionMessage

	for _, message := range *messages {
		if message.Role != "system" || strings.Contains(message.Content, "#function_call") {
			otherMessages = append(otherMessages, message)
		}
	}

	*messages = otherMessages
}
