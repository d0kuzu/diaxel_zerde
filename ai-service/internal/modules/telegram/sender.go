package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type TelegramMessageReq struct {
	ChatID    int64  `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

type Sender struct {
	globalTicker *time.Ticker
	chatLastSent map[int64]time.Time
	mu           sync.Mutex
	client       *http.Client
}

func NewSender() *Sender {
	return &Sender{
		globalTicker: time.NewTicker(35 * time.Millisecond),
		chatLastSent: make(map[int64]time.Time),
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *Sender) Send(botToken string, chatID int64, text string) error {
	const limit = 4096
	runes := []rune(text)

	for i := 0; i < len(runes); i += limit {
		end := i + limit
		if end > len(runes) {
			end = len(runes)
		}

		err := s.sendPart(botToken, chatID, string(runes[i:end]))
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Sender) sendPart(botToken string, chatID int64, text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	reqBody := TelegramMessageReq{
		ChatID:    chatID,
		Text:      text,
		ParseMode: "Markdown",
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	maxRetries := 3
	for attempt := 0; attempt < maxRetries; attempt++ {
		<-s.globalTicker.C

		s.mu.Lock()
		lastSent, exists := s.chatLastSent[chatID]
		if exists {
			elapsed := time.Since(lastSent)
			if elapsed < time.Second {
				time.Sleep(time.Second - elapsed)
			}
		}

		s.chatLastSent[chatID] = time.Now()
		s.mu.Unlock()

		resp, err := s.client.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return fmt.Errorf("failed to send request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			return nil
		}

		if resp.StatusCode == http.StatusTooManyRequests {
			retryAfter := 1
			if retryHeader := resp.Header.Get("Retry-After"); retryHeader != "" {
				fmt.Sscanf(retryHeader, "%d", &retryAfter)
			}
			log.Printf("[Telegram Sender] Rate limited (429)! Retrying after %d seconds...", retryAfter)
			time.Sleep(time.Duration(retryAfter) * time.Second)
			continue
		}

		return fmt.Errorf("telegram API returned status: %d", resp.StatusCode)
	}

	return fmt.Errorf("failed to send message after %d attempts", maxRetries)
}
