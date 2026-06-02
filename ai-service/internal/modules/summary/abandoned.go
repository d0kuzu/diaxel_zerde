package summary

import (
	"context"
	"fmt"
	"log"
	"time"

	"diaxel/internal/constants"
	"diaxel/internal/grpc/db"
	"diaxel/internal/modules/campuslogin"
	"diaxel/internal/modules/llm"
	
	"github.com/sashabaranov/go-openai"
)

type AbandonedSummarizer struct {
	dbClient     *db.Client
	campusClient *campuslogin.Client
	llmClient    *llm.Client
}

func NewAbandonedSummarizer(dbClient *db.Client, campusClient *campuslogin.Client, llmClient *llm.Client) *AbandonedSummarizer {
	return &AbandonedSummarizer{
		dbClient:     dbClient,
		campusClient: campusClient,
		llmClient:    llmClient,
	}
}

func (s *AbandonedSummarizer) Start(ctx context.Context) {
	// Calculate time until next 7:30 AM in America/Winnipeg
	loc, err := time.LoadLocation("America/Winnipeg")
	if err != nil {
		log.Printf("[AbandonedSummarizer] Error loading timezone: %v\n", err)
		return
	}

	for {
		now := time.Now().In(loc)
		next := time.Date(now.Year(), now.Month(), now.Day(), 7, 30, 0, 0, loc)
		if now.After(next) {
			next = next.Add(24 * time.Hour)
		}
		duration := next.Sub(now)

		select {
		case <-ctx.Done():
			log.Println("[AbandonedSummarizer] Stopped.")
			return
		case <-time.After(duration):
			s.processAbandonedChats(ctx)
		}
	}
}

func (s *AbandonedSummarizer) processAbandonedChats(ctx context.Context) {
	log.Println("[AbandonedSummarizer] Starting daily summarization...")
	
	chats, err := s.dbClient.GetUnreviewedActiveChats()
	if err != nil {
		log.Printf("[AbandonedSummarizer] Error getting unreviewed active chats: %v", err)
		return
	}

	for _, chat := range chats {
		if chat.CustomerId == "" {
			continue
		}

		updatedAt, err := time.Parse(time.RFC3339, chat.UpdatedAt)
		if err != nil {
			log.Printf("[AbandonedSummarizer] Error parsing UpdatedAt '%s' for chat %s: %v", chat.UpdatedAt, chat.Id, err)
			continue
		}

		if time.Since(updatedAt) < 1*time.Hour {
			continue // Chat is still active or updated recently
		}

		messages, err := s.dbClient.GetAllChatMessages(chat.Id)
		if err != nil {
			log.Printf("[AbandonedSummarizer] Error getting messages for chat %s: %v", chat.Id, err)
			continue
		}

		if len(messages) == 0 {
			continue
		}

		hasUserMessage := false
		var userMessagesText string
		for _, msg := range messages {
			if msg.Role == "user" {
				hasUserMessage = true
			}
			userMessagesText += fmt.Sprintf("%s: %s\n", msg.Role, msg.Content)
		}

		var summary string
		if !hasUserMessage {
			summary = "Lead didn't answer on the first message."
		} else {
			// Ask LLM for summary
			prompt := []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: constants.AbandonedChatSummaryPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userMessagesText,
				},
			}
			response, err := s.llmClient.GetAnswer(ctx, prompt)
			if err != nil || len(response.Choices) == 0 {
				log.Printf("[AbandonedSummarizer] Error getting LLM answer for chat %s: %v", chat.Id, err)
				continue
			}
			summary = response.Choices[0].Message.Content
		}

		summary += fmt.Sprintf("\n\nFollowup Stage: %d", chat.FollowupStage)

		campusRecord, err := s.dbClient.GetCampusloginByUserId(chat.CustomerId)
		var contactID int
		if err != nil {
			log.Printf("[AbandonedSummarizer] Failed to get ContactID for user %s: %v", chat.CustomerId, err)
			contactID = 5972449 // default fallback
		} else {
			contactID = int(campusRecord.ContactId)
		}

		err = s.campusClient.AddNewNote(ctx, contactID, summary)
		if err != nil {
			log.Printf("[AbandonedSummarizer] Failed to add note for chat %s: %v", chat.Id, err)
			continue
		}

		_, err = s.dbClient.UpdateChatIsReviewed(chat.Id, true)
		if err != nil {
			log.Printf("[AbandonedSummarizer] Failed to update chat is_reviewed for %s: %v", chat.Id, err)
		} else {
			log.Printf("[AbandonedSummarizer] Successfully summarized abandoned chat for user %s", chat.CustomerId)
		}
	}

	log.Println("[AbandonedSummarizer] Finished daily summarization.")
}
