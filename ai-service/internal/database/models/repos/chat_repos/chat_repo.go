package chat_repos

import (
	"diaxel/internal/database"
	. "diaxel/internal/database/models"
	"errors"
	"gorm.io/gorm"
)

func CheckIfExist(userId string) (Chat, error) {
	db := database.GetDB()
	var chat Chat

	if err := db.Preload("Messages", func(db *gorm.DB) *gorm.DB {
		return db.Order("id ASC")
	}).Where("user_id = ?", userId).First(&chat).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Chat{}, nil
		}
		return Chat{}, err
	}

	return chat, nil
}

func Save(userId string, messages []Message) error {
	db := database.GetDB()

	var chat Chat
	result := db.Preload("Messages").First(&chat, "user_id = ?", userId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			chat = Chat{
				UserID:   userId,
				Messages: messages,
			}
			if err := db.Create(&chat).Error; err != nil {
				return errors.New("error creating the record: " + err.Error())
			}
			return nil
		}
		return errors.New("error checking the record: " + result.Error.Error())
	}

	newMessages := messages[len(chat.Messages):]

	if len(newMessages) > 0 {
		if err := db.Create(&newMessages).Error; err != nil {
			return errors.New("error adding new messages: " + err.Error())
		}
	}

	return nil
}

func GetAll() ([]Chat, error) {
	db := database.GetDB()
	var chats []Chat

	if err := db.Table("chats").Find(&chats).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []Chat{}, nil
		}
		return []Chat{}, err
	}
	return chats, nil
}

func SetClientStatusTrue(userID string) error {
	db := database.GetDB()

	return db.Model(&Chat{}).Where("user_id = ?", userID).Update("is_client", true).Error
}

func ClearMessages() error {
	db := database.GetDB()

	if err := db.Exec("TRUNCATE TABLE messages").Error; err != nil {
		return err
	}

	return nil
}

func ClearChatMessages(userId string) error {
	db := database.GetDB()

	return db.Where("chat_user_id = ?", userId).Delete(&Message{}).Error
}
