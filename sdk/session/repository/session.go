package repository

import (
	"github.com/Fiddler25/ddd-sample-app/sdk/session/model"
	"gorm.io/gorm"
	"log"
)

type SessionRepository struct {
	db *gorm.DB
}

func NewSession(db *gorm.DB) SessionRepository {
	return SessionRepository{db: db}
}

func (r SessionRepository) Create(sess *model.Session) *model.Session {
	if err := r.db.Create(&sess).Error; err != nil {
		log.Fatalln(err)
	}

	return sess
}
