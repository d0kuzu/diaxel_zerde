package twilio

import (
	"context"
	"fmt"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func (c *Client) SendMessage(ctx context.Context, to, message string) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(c.twilioFromNumber)
	params.SetBody(message)

	resp, err := c.client.Api.CreateMessage(params)
	if err != nil {
		return err
	}

	fmt.Println("Message SID:", *resp.Sid)
	return nil
}
