package user

import (
	"gorm.io/gorm"
	"log"
)

type GetRepository struct {
	db *gorm.DB
}

func NewGetRepository(db *gorm.DB) GetRepository {
	return GetRepository{db: db}
}

func (r GetRepository) ByUserID(userID int) User {
	var ret User
	if err := r.db.Find(&ret, userID).Error; err != nil {
		log.Fatal(err)
	}

	return ret
}
