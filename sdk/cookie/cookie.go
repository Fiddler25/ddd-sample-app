package cookie

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type CookieAttributes struct {
	Name  string
	Value string
}

const (
	path    = "/"
	expires = time.Hour * 24 * 365 * 20 // 20年
)

func Set(c echo.Context, userID string, token string) {
	bc := baseCookie()

	setUserID(c, bc, userID)
	setRememberToken(c, bc, token)
}

func baseCookie() *http.Cookie {
	return &http.Cookie{
		Path:    path,
		Expires: time.Now().Add(expires),
	}
}

func setUserID(c echo.Context, cookie *http.Cookie, userID string) {
	cookie.Name = "user_id"
	cookie.Value = userID

	c.SetCookie(cookie)
}

func setRememberToken(c echo.Context, cookie *http.Cookie, token string) {
	cookie.Name = "remember_token"
	cookie.Value = token

	c.SetCookie(cookie)
}
