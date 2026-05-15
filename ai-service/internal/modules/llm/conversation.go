package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/sashabaranov/go-openai"
)

const defaultTimezone = "America/Winnipeg"

func (c *Client) Conversation(ctx context.Context, userId, assistantId, userMessage string) (string, error) {
	log.Printf("Сообщение от пользователя %s: %s", userId, userMessage)
	messages, err := c.GetMessages(assistantId, userId)
	if err != nil {
		return "", err
	}

	c.AddMessage(&messages, "user", userMessage)

	response, err := c.GetAnswer(ctx, messages)
	if err != nil {
		return "", err
	}

	// Tool call loop: keep processing until AI returns a text response
	for len(response.Choices) > 0 && len(response.Choices[0].Message.ToolCalls) > 0 {
		// Add the assistant message with tool calls to the conversation
		assistantMsg := response.Choices[0].Message
		// go-openai uses `json:"content,omitempty"` — empty string is dropped to null in JSON.
		// OpenAI API rejects null content on assistant messages, so we must ensure it's non-empty.
		if assistantMsg.Content == "" && len(assistantMsg.ToolCalls) > 0 {
			assistantMsg.Content = " "
		}
		messages = append(messages, assistantMsg)

		// Process each tool call
		for _, toolCall := range assistantMsg.ToolCalls {
			log.Printf("Функция вызвана: %s", toolCall.Function.Name)
			log.Printf("Аргументы: %s", toolCall.Function.Arguments)

			result, err := c.executeFunction(ctx, toolCall.Function.Name, toolCall.Function.Arguments)
			if err != nil {
				log.Printf("Ошибка выполнения функции %s: %v", toolCall.Function.Name, err)
				result = fmt.Sprintf("Error executing function: %s", err.Error())
			}

			log.Printf("Результат функции %s: %s", toolCall.Function.Name, result)

			// Add tool response message
			messages = append(messages, openai.ChatCompletionMessage{
				Role:       openai.ChatMessageRoleTool,
				Content:    result,
				ToolCallID: toolCall.ID,
			})
		}

		// Make another request to OpenAI so it can generate a text response based on the function results
		response, err = c.GetAnswer(ctx, messages)
		if err != nil {
			return "", err
		}
	}

	assistantResponse := response.Choices[0].Message.Content
	log.Printf("Ответ пользователю %s от ИИ: %s\n", userId, assistantResponse)

	c.AddMessage(&messages, "assistant", assistantResponse)

	err = c.SaveMessages(assistantId, userId, messages)
	if err != nil {
		return "", err
	}

	log.Println("Конец запроса")
	return assistantResponse, nil
}

// executeFunction dispatches tool calls to the appropriate handler
func (c *Client) executeFunction(ctx context.Context, functionName, argsJSON string) (string, error) {
	switch functionName {
	case "get_available_slots":
		return c.handleGetAvailableSlots(ctx, argsJSON)
	case "create_booking":
		return c.handleCreateBooking(ctx, argsJSON)
	default:
		return "", fmt.Errorf("unknown function: %s", functionName)
	}
}

// handleGetAvailableSlots processes the get_available_slots function call
func (c *Client) handleGetAvailableSlots(ctx context.Context, argsJSON string) (string, error) {
	var args struct {
		Date string `json:"date"`
	}
	if err := json.Unmarshal([]byte(argsJSON), &args); err != nil {
		return "", fmt.Errorf("failed to parse arguments: %w", err)
	}

	if args.Date == "" {
		return "Error: date parameter is required in YYYY-MM-DD format", nil
	}

	slots, err := c.calcom.GetAvailableSlots(ctx, args.Date, defaultTimezone)
	if err != nil {
		return "", fmt.Errorf("failed to get slots from Cal.com: %w", err)
	}

	if len(slots) == 0 {
		return fmt.Sprintf("No available slots found for %s. Please suggest another date.", args.Date), nil
	}

	return fmt.Sprintf("Available time slots for %s:\n%s", args.Date, strings.Join(slots, "\n")), nil
}

// handleCreateBooking processes the create_booking function call
func (c *Client) handleCreateBooking(ctx context.Context, argsJSON string) (string, error) {
	var args struct {
		StartTime     string `json:"start_time"`
		AttendeeName  string `json:"attendee_name"`
		AttendeeEmail string `json:"attendee_email"`
	}
	if err := json.Unmarshal([]byte(argsJSON), &args); err != nil {
		return "", fmt.Errorf("failed to parse arguments: %w", err)
	}

	if args.StartTime == "" || args.AttendeeName == "" || args.AttendeeEmail == "" {
		return "Error: start_time, attendee_name, and attendee_email are all required", nil
	}

	booking, err := c.calcom.CreateBooking(ctx, args.StartTime, args.AttendeeName, args.AttendeeEmail, defaultTimezone)
	if err != nil {
		return "", fmt.Errorf("failed to create booking on Cal.com: %w", err)
	}

	return fmt.Sprintf("Booking confirmed! Details:\n- ID: %d\n- Title: %s\n- Status: %s\n- Start: %s\n- End: %s",
		booking.ID, booking.Title, booking.Status, booking.Start, booking.End), nil
}
