package auth

import (
	"github.com/Fiddler25/sample-app/db"
	"github.com/Fiddler25/sample-app/domain/user"
	"github.com/Fiddler25/sample-app/sdk/cookie"
	"github.com/Fiddler25/sample-app/sdk/session"
	"github.com/Fiddler25/sample-app/sdk/validator"
	"github.com/Fiddler25/sample-app/usecase/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SignUp(c echo.Context) error {
	var req auth.SignUpRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	validate := validator.Validate()
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := validate.VarWithValue(req.Password, req.PasswordConfirmation, "eqfield"); err != nil {
		return c.JSON(http.StatusBadRequest, "パスワードが一致しません")
	}

	if res, err := auth.NewSignUpUsecase(db.Conn()).Execute(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, res)
	}
}

func Login(c echo.Context) error {
	var req auth.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := auth.NewLoginUsecase(db.Conn()).Execute(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	if err := startSession(c, res.ID); err != nil {
		return c.JSON(http.StatusUnauthorized, "ログインできません。")
	}

	return c.JSON(http.StatusOK, res)
}

func startSession(c echo.Context, userID user.UserID) error {
	var oldSessionID session.ID
	if ck, err := c.Cookie(cookie.CookieNameSession); err != nil {
		oldSessionID = ""
	} else {
		oldSessionID = session.ID(ck.Value)
	}

	if ck, err := session.NewService(db.Conn()).Start(oldSessionID, userID); err != nil {
		return err
	} else {
		c.SetCookie(ck)
		return nil
	}
}
