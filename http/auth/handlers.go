package auth

import (
	"github.com/Fiddler25/ddd-sample-app/gorm"
	"github.com/Fiddler25/ddd-sample-app/usecase/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Login(c echo.Context) error {
	var req auth.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	res := auth.NewLoginUsecase(gorm.DB()).Execute(req)
	if res == nil {
		return c.JSON(http.StatusUnauthorized, "ログインできません。")
	}

	return c.JSON(http.StatusOK, res)
}
