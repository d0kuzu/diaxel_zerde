package repository

import (
	"diaxel_zerde/database-service/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CampusloginRepository struct {
	db *gorm.DB
}

func NewCampusloginRepository(db *gorm.DB) *CampusloginRepository {
	return &CampusloginRepository{db: db}
}

func (r *CampusloginRepository) GetByUserId(userId string) (*models.Campuslogin, error) {
	var campuslogin models.Campuslogin
	err := r.db.Where("user_id = ?", userId).First(&campuslogin).Error
	if err != nil {
		return nil, err
	}
	return &campuslogin, nil
}

func (r *CampusloginRepository) Upsert(campuslogin *models.Campuslogin) error {
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"contact_id", "program_id"}),
	}).Create(campuslogin).Error
}
