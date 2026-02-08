package twilio

import (
	"github.com/twilio/twilio-go"
)

type Client struct {
	client           *twilio.RestClient
	twilioFromNumber string
}

func InitClient(TwilioAccountSID, TwilioAuthToken string) *Client {
	return &Client{
		client: twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: TwilioAccountSID,
			Password: TwilioAuthToken,
		}),
		twilioFromNumber: "",
	}
}
