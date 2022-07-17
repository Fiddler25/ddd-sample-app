package auth

import (
	"github.com/Fiddler25/ddd-sample-app/domain/model"
	"github.com/Fiddler25/ddd-sample-app/gorm"
	"github.com/Fiddler25/ddd-sample-app/sdk/cookie"
	"github.com/Fiddler25/ddd-sample-app/sdk/session"
	"github.com/Fiddler25/ddd-sample-app/usecase/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Login(c echo.Context) error {
	var req auth.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	res := auth.NewLoginUsecase(gorm.DB()).Execute(c, req)
	if res == nil {
		return c.JSON(http.StatusUnauthorized, "ログインできません。")
	}

	if err := startSession(c, res.ID); err != nil {
		return c.JSON(http.StatusUnauthorized, "ログインできません。")
	}

	return c.JSON(http.StatusOK, res)
}

func startSession(c echo.Context, userID model.UserID) error {
	var oldSessionID session.ID
	if ck, err := c.Cookie(cookie.CookieNameSession); err != nil {
		oldSessionID = ""
	} else {
		oldSessionID = session.ID(ck.Value)
	}

	sRepo := session.NewRepository(gorm.DB())
	if oldSessionID != "" {
		if oldSession, err := sRepo.Find(oldSessionID); err != nil {
			return err
		} else {
			sRepo.Delete(oldSession)
		}
	}

	sessionID := session.NewID()
	sess := session.NewModel(sessionID, userID)
	sRepo.Create(sess)

	cookie.Set(c, string(sess.SessionID))

	return nil
}
