package auth

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) GetByEmail(email Email) (*User, error) {
	var ret *User
	if err := r.db.Where("email = ?", email).Take(&ret).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	return ret, nil
}

func (r Repository) Create(ret *User) *User {
	if err := r.db.Create(&ret).Error; err != nil {
		log.Fatal(err)
	}
	return ret
}
