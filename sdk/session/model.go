package session

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/sdk/rand"
	"log"
)

type ID string

type Session struct {
	SessionID ID
	UserID    model.UserID
}

const sessionIDLength = 32

func NewModel(sessID ID, userID model.UserID) *Session {
	return &Session{
		SessionID: sessID,
		UserID:    userID,
	}
}

func NewID() ID {
	sessID, err := rand.GenerateRandomString(sessionIDLength)
	if err != nil {
		log.Fatal(err)
	}

	return ID(sessID)
}
