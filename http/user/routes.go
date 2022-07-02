package user

import "github.com/labstack/echo/v4"

func Setup(e *echo.Echo) {
	e.GET("api/users/:user_id", Get)
}
