package campuslogin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"diaxel/internal/constants"
)

type AddNewNoteRequest struct {
	OrgId       int    `json:"orgId"`
	MailListId  int    `json:"mailListId"`
	ContactId   int    `json:"contactId"`
	ContactType string `json:"contactType"`
	StaffId     int    `json:"staffId"`
	NoteText    string `json:"noteText"`
	NoteOptions string `json:"noteOptions"`
	Type        string `json:"type"`
	Association any    `json:"association"`
	CreateDate  string `json:"createDate"`
	CreatedBy   any    `json:"createdBy"`
}

func (c *Client) AddNewNote(ctx context.Context, contactId int, summary string) error {
	loc, err := time.LoadLocation(constants.DefaultTimezone)
	if err != nil {
		return fmt.Errorf("failed to load timezone %s: %w", constants.DefaultTimezone, err)
	}
	createDate := time.Now().In(loc).Format("2006-01-02T15:04:05.000Z")

	reqBody := AddNewNoteRequest{
		OrgId:       24900,
		MailListId:  24901001,
		ContactId:   contactId,
		ContactType: "",
		StaffId:     1,
		NoteText:    summary,
		NoteOptions: "",
		Type:        "",
		Association: nil,
		CreateDate:  createDate,
		CreatedBy:   nil,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://connectorapi.campuslogin.com/api/Notes/AddNewNote", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	// Using the provided API key
	req.Header.Set("X-Api-Key", c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
