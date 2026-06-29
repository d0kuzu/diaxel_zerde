package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"diaxel/internal/constants"
)

func (c *Client) executeFunction(ctx context.Context, functionName, argsJSON, userId, assistantId string) (string, error) {
	switch functionName {
	case "calcom_get_available_slots":
		return c.handleGetAvailableSlots(ctx, argsJSON)
	case "calcom_create_booking":
		return c.handleCreateBooking(ctx, argsJSON)
	default:
		return "", fmt.Errorf("unknown function: %s", functionName)
	}
}

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

	slots, err := c.calcom.GetAvailableSlots(ctx, args.Date, constants.DefaultTimezone)
	if err != nil {
		return "", fmt.Errorf("failed to get slots from Cal.com: %w", err)
	}

	if len(slots) == 0 {
		return fmt.Sprintf("No available slots found for %s. Please suggest another date.", args.Date), nil
	}

	return fmt.Sprintf("Available time slots for %s:\n%s", args.Date, strings.Join(slots, "\n")), nil
}

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

	booking, err := c.calcom.CreateBooking(ctx, args.StartTime, args.AttendeeName, args.AttendeeEmail, constants.DefaultTimezone)
	if err != nil {
		return "", fmt.Errorf("failed to create booking on Cal.com: %w", err)
	}

	return fmt.Sprintf("Booking confirmed! Details:\n- ID: %d\n- Title: %s\n- Status: %s\n- Start: %s\n- End: %s",
		booking.ID, booking.Title, booking.Status, booking.Start, booking.End), nil
}
