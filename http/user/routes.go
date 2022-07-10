package user

import "github.com/labstack/echo/v4"

func Setup(e *echo.Echo) {
	e.GET("/api/users/:user_id", Get)
	e.POST("/api/users", Create)
	e.PUT("/api/users/:user_id", Update)
	e.DELETE("/api/users/:user_id", Delete)
}
