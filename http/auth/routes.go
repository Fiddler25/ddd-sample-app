package auth

import "github.com/labstack/echo/v4"

func Setup(g *echo.Group) {
	g.POST("/login", Login)
}
