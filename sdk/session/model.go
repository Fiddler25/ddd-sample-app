package session

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/sdk/rand"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
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

func Login(c echo.Context, userID model.UserID) {
	sess, _ := session.Get("session", c)
	sess.Values["user_id"] = userID
	sess.Save(c.Request(), c.Response())
}
