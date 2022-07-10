package user

import "github.com/labstack/echo/v4"

func Setup(g *echo.Group) {
	g.GET("/users/:user_id", Get)
	g.POST("/users", Create)
	g.PUT("/users/:user_id", Update)
	g.DELETE("/users/:user_id", Delete)
}
