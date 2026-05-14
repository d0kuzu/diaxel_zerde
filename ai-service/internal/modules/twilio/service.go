package twilio

import (
	"context"
	"fmt"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func (c *Client) SendMessage(ctx context.Context, accountSID, authToken, from, to, message string) error {
	client := c.GetRestClient(accountSID, authToken)

	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		return err
	}

	if resp.Sid != nil {
		fmt.Printf("[Twilio REST] Message sent successfully. SID: %s, From: %s, To: %s\n", *resp.Sid, from, to)
	} else {
		fmt.Printf("[Twilio REST] Message sent, but SID is nil. From: %s, To: %s\n", from, to)
	}
	return nil
}
