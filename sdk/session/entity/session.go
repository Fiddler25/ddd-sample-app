package entity

import (
	"github.com/Fiddler25/ddd-sample-app/sdk/rand"
	"github.com/Fiddler25/ddd-sample-app/sdk/session/model"
	"log"
)

const sessionIDLength = 32

func NewSessionID() model.SessionID {
	sessID, err := rand.GenerateRandomString(sessionIDLength)
	if err != nil {
		log.Fatal(err)
	}

	return model.SessionID(sessID)
}
