package twilio

import (
	"diaxel/internal/config"
	"diaxel/internal/modules/llm"
	twilio2 "diaxel/internal/modules/twilio"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TwilioWebhookHandler struct {
	cfg    *config.Settings
	LLM    *llm.Client
	twilio *twilio2.Client
}

func NewTwilioWebhookHandler(cfg *config.Settings, llmClient *llm.Client, twilioClient *twilio2.Client) *TwilioWebhookHandler {
	return &TwilioWebhookHandler{cfg: cfg, LLM: llmClient, twilio: twilioClient}
}

func (h *TwilioWebhookHandler) HandleWebhook(c *gin.Context) {
	from := c.PostForm("From")
	body := c.PostForm("Body")

	answer, err := h.LLM.Conversation(c, from, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.twilio.SendMessage(c, from, answer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.XML(200, gin.H{
		"Response": gin.H{
			"Message": "OK",
		},
	})
}
