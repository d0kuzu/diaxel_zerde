package llm

import (
	"diaxel/internal/constants"
	. "diaxel/internal/database/models"
	"diaxel/internal/database/models/repos/chat_repos"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
)

func GetHistory(userId string) ([]Message, error) {
	chat, err := chat_repos.CheckIfExist(userId)

	return chat.Messages, err
}

func GetAllChats() ([]string, error) {
	var parsedChats []string

	chats, err := chat_repos.GetAll()
	if err == nil && len(chats) > 0 {
		for _, chat := range chats {

			parsedChats = append(parsedChats, chat.UserID.String())
		}

		return parsedChats, err
	}

	return parsedChats, err
}

func GetMessages(userId string) ([]openai.ChatCompletionMessage, error) {
	chat, err := chat_repos.CheckIfExist(userId)
	var messages []openai.ChatCompletionMessage
	rawMessages := chat.Messages

	if err != nil {
		return messages, err
	} else if len(chat.Messages) != 0 {
		convertedMessages, err := ConvertToOpenaiMessage(rawMessages)
		if err != nil {
			return messages, err
		}

		messages = append(messages, convertedMessages...)
		// CheckSystemMessages(&messages)
	} //else if len(chat.Messages) == 0 {
	//	messages = append(messages, config.StartBotMessage...)
	//}
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
	RemoveSystemMessages(&messages)
	arrayMessages := ConvertToMessage(userId, messages)

	err := chat_repos.Save(userId, arrayMessages)
	if err != nil {
		return err
	}

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
