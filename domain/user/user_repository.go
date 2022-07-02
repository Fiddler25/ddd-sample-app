package user

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) ByUserID(userID int) User {
	var ret User
	if err := r.db.Find(&ret, userID).Error; err != nil {
		log.Fatal(err)
	}

	return ret
}
