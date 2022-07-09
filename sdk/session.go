package sdk

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context, userID int) {
	sess, _ := session.Get("session", c)
	sess.Values["user_id"] = userID
	sess.Save(c.Request(), c.Response())
}
