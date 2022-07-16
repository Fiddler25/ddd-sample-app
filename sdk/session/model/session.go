package model

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type SessionID string

type Session struct {
	SessionID SessionID
	UserID    model.UserID
}

func NewSession(sessID SessionID, userID model.UserID) *Session {
	return &Session{
		SessionID: sessID,
		UserID:    userID,
	}
}

func Login(c echo.Context, userID model.UserID) {
	sess, _ := session.Get("session", c)
	sess.Values["user_id"] = userID
	sess.Save(c.Request(), c.Response())
}
