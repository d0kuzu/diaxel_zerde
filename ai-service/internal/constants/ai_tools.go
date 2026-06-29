package constants

import "github.com/sashabaranov/go-openai"

var Tools = []openai.Tool{
	{
		Type: openai.ToolTypeFunction,
		Function: &openai.FunctionDefinition{
			Name:        "calcom_get_available_slots",
			Description: "Get available appointment time slots for a specific date from Cal.com. Call this when the user mentions a day or date they want to visit.",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"date": map[string]interface{}{
						"type":        "string",
						"description": "The date to check availability for, in YYYY-MM-DD format (e.g. 2026-05-16)",
					},
				},
				"required": []string{"date"},
			},
		},
	},
	{
		Type: openai.ToolTypeFunction,
		Function: &openai.FunctionDefinition{
			Name:        "calcom_create_booking",
			Description: "Create a booking appointment at a specific time slot on Cal.com. Call this only after the user has selected an available time slot and provided their name and email address.",
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"start_time": map[string]interface{}{
						"type":        "string",
						"description": "The start time of the appointment in ISO 8601 UTC format (e.g. 2026-05-16T14:00:00Z)",
					},
					"attendee_name": map[string]interface{}{
						"type":        "string",
						"description": "Full name of the person booking the appointment",
					},
					"attendee_email": map[string]interface{}{
						"type":        "string",
						"description": "Email address of the person booking the appointment",
					},
				},
				"required": []string{"start_time", "attendee_name", "attendee_email"},
			},
		},
	},
}
