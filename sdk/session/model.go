package session

import (
	"github.com/Fiddler25/ddd-sample-app/domain/user"
	"github.com/Fiddler25/ddd-sample-app/sdk/rand"
)

type ID string

type Session struct {
	SessionID ID
	UserID    user.UserID
}

const sessionIDLength = 32

func NewSession(sessID ID, userID user.UserID) *Session {
	return &Session{
		SessionID: sessID,
		UserID:    userID,
	}
}

func NewID() ID {
	sessID, _ := rand.GenerateRandomString(sessionIDLength)
	return ID(sessID)
}
