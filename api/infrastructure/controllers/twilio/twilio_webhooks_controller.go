package twilio

import (
	"context"
	"diaxel/config"
	"diaxel/services/llm"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type TwilioWebhookHandler struct {
	cfg *config.Settings
	LLM *llm.Client
}

func NewTwilioWebhookHandler(cfg *config.Settings, llmClient *llm.Client) *TwilioWebhookHandler {
	return &TwilioWebhookHandler{cfg: cfg, LLM: llmClient}
}

func (h *TwilioWebhookHandler) HandleWebhook(c *gin.Context) {
	from := c.PostForm("From")
	body := c.PostForm("Body")

	// Генерация ответа через LLM (если нужно)
	reply := "Спасибо за ваше сообщение!"

	// Отправляем ответ пользователю
	err := h.SendMessage(c, from, reply)
	if err != nil {
		log.Println("Twilio send error:", err)
	}

	// Twilio требует TwiML-ответ
	c.XML(200, gin.H{
		"Response": gin.H{
			"Message": "OK",
		},
	})
}

func (h *TwilioWebhookHandler) SendMessage(ctx context.Context, to, message string) error {
	accountSID := h.cfg.TwilioAccountSID
	authToken := h.cfg.TwilioAuthToken
	from := "" //TODO: номер бота

	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSID + "/Messages.json"

	data := url.Values{}
	data.Set("To", to)
	data.Set("From", from)
	data.Set("Body", message)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("twilio error: status=%d body=%s", resp.StatusCode, string(body))
	}

	return nil
}
