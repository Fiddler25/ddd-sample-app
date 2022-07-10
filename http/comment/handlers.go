package comment

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Create(c echo.Context) error {
	return c.JSON(http.StatusCreated, "success")
}
