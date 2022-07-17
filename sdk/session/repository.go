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

func (r Repository) Find(sessionID ID) (*Session, error) {
	var session Session
	if err := r.db.Where("session_id = ?", sessionID).Take(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (r Repository) Delete(sess *Session) {
	if err := r.db.
		Where("session_id = ?", sess.SessionID).
		Where("user_id", sess.UserID).
		Delete(sess).
		Error; err != nil {
		log.Fatalln(err)
	}
}
