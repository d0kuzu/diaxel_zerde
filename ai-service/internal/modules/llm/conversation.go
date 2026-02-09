package llm

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (c *Client) Conversation(ctx *gin.Context, userId string, userMessage string) (string, error) {
	log.Printf("Сообщение от пользователя %s: %s", userId, userMessage)
	messages, err := GetMessages(userId)
	if err != nil {
		return "", err
	}

	AddMessage(&messages, "user", userMessage)

	response, err := c.GetAnswer(ctx, messages)
	if err != nil {
		return "", err
	}

	toolCalls := response.Choices[0].Message.ToolCalls
	if len(toolCalls) > 0 {
		argsJSON := toolCalls[0].Function.Arguments
		log.Printf("Имя функции - %s", toolCalls[0].Function.Name)
		log.Printf("Данные запроса пользователя %s: %s\n", userId, argsJSON)

		switch toolCalls[0].Function.Name {
		case "bookcampussanjose", "bookcampussanfrancisco":
			log.Printf("____________BOOKING_________")
			//makeClient := make.New()
			//bookUrl, err := makeClient.SanJoseBook(ctx, constants.SanJoseBookingWebhook)
			//if err != nil {
			//	return "", err
			//}
			//
			//answer := fmt.Sprintf("your appointment link: %s", bookUrl)
			//
			//AddMessage(&messages, "assistant", answer)
		}
	}
	log.Printf("Ответ пользователю %s от ИИ: %s\n", userId, response.Choices[0].Message.Content)

	AddMessage(&messages, "assistant", response.Choices[0].Message.Content)

	err = SaveMessages(userId, messages)
	if err != nil {
		return "", err
	}

	log.Println("Конец запроса")
	return response.Choices[0].Message.Content, nil
}
