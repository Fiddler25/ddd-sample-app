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

func Set(c echo.Context, token string) {
	bc := baseCookie()

	setRememberToken(c, bc, token)
}

func baseCookie() *http.Cookie {
	return &http.Cookie{
		Path:    path,
		Expires: time.Now().Add(expires),
	}
}

func setRememberToken(c echo.Context, cookie *http.Cookie, token string) {
	cookie.Name = "remember_token"
	cookie.Value = token

	c.SetCookie(cookie)
}
