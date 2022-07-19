package auth

import "github.com/labstack/echo/v4"

func Setup(g *echo.Group) {
	g.POST("/signup", SignUp)
	g.POST("/login", Login)
}
