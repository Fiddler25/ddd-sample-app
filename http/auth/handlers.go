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

	auth.NewLoginUsecase(gorm.DB())

	return c.JSON(http.StatusOK, req)
}
