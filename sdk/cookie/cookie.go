package cookie

import (
	"net/http"
	"time"
)

const (
	CookieNameSession = "session_sample"
	path              = "/"
	expires           = time.Hour * 24 * 365 * 20 // 20年
)

func Create(sessionID string) *http.Cookie {
	return &http.Cookie{
		Name:    CookieNameSession,
		Value:   sessionID,
		Path:    path,
		Expires: time.Now().Add(expires),
	}
}
