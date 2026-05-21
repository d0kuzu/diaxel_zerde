package models

type Campuslogin struct {
	UserId    string `gorm:"primaryKey;column:user_id"`
	ContactID int    `gorm:"column:contact_id"`
	ProgramID int    `gorm:"column:program_id"`
}
