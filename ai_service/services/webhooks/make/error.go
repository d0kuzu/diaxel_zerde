package make

import "net/http"

type WebhookError struct {
	Status int
	Body   string
}

func (e *WebhookError) Error() string {
	return "webhook request failed with status " + http.StatusText(e.Status)
}
