package comment

import "github.com/labstack/echo/v4"

func Setup(g *echo.Group) {
	g.POST("/comments", Create)
}
