package constants

import "github.com/sashabaranov/go-openai"

var Tools = []openai.Tool{
	{
		Type: openai.ToolTypeFunction,
		Function: &openai.FunctionDefinition{
			Name:        "bookcampussanjose",
			Description: "Use this function only once and only after the customer has explicitly confirmed their interest in booking a campus tour in San Jose.\n\nOnce the confirmation is received, trigger the function to return the appointment booking link.",
			Parameters: map[string]interface{}{
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		},
	},
	{
		Type: openai.ToolTypeFunction,
		Function: &openai.FunctionDefinition{
			Name:        "bookcampussanfrancisco",
			Description: "Use this function only once and only after the customer has explicitly confirmed their interest in booking a campus tour in San Francisco.\n\nOnce the confirmation is received, trigger the function to return the appointment booking link.",
			Parameters: map[string]interface{}{
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		},
	},
}
