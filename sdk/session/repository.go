package session

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) Create(sess *Session) *Session {
	if err := r.db.Create(&sess).Error; err != nil {
		log.Fatalln(err)
	}

	return sess
}
