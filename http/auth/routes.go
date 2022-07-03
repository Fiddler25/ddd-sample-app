package auth

import "github.com/labstack/echo/v4"

func Setup(e *echo.Echo) {
	e.POST("/api/login", Login)
}
