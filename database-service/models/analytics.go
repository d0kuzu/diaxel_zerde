package models

// Analytics - соответствует таблице analytics в БД
type Analytics struct {
	ID             string  `json:"id" db:"id"`
	AssistantID    string  `json:"assistant_id" db:"assistant_id"`
	TotalChats     int32   `json:"total_chats" db:"total_chats"`
	ActiveUsers    int32   `json:"active_users" db:"active_users"`
	EngagementRate float64 `json:"engagement_rate" db:"engagement_rate"`
}

// TableName returns the table name for Analytics model
func (Analytics) TableName() string {
	return "analytics"
}
