package cookie

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const (
	CookieNameSession = "session_sample"
	path              = "/"
	expires           = time.Hour * 24 * 365 * 20 // 20年
)

func Set(c echo.Context, sessionID string) {
	cookie := &http.Cookie{
		Name:    CookieNameSession,
		Value:   sessionID,
		Path:    path,
		Expires: time.Now().Add(expires),
	}

	c.SetCookie(cookie)
}
