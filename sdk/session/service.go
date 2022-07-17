package session

import (
	"github.com/Fiddler25/ddd-sample-app/domain/user"
	"github.com/Fiddler25/ddd-sample-app/sdk/cookie"
	"gorm.io/gorm"
	"net/http"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return Service{db: db}
}

func (s Service) Start(oldSessionID ID, userID user.UserID) (*http.Cookie, error) {
	tx := s.db.Begin()

	sRepo := NewRepository(tx)
	if oldSessionID != "" {
		if _, err := sRepo.Find(oldSessionID); err != nil {
			return nil, err
		} else {
			sRepo.Delete(&Session{
				SessionID: oldSessionID, UserID: userID,
			})
		}
	}

	sessionID := NewID()
	sess := sRepo.Create(NewSession(sessionID, userID))

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return cookie.Create(string(sess.SessionID)), nil
}
