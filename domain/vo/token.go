package vo

import (
	"github.com/labstack/gommon/random"
)

type Token string

func NewRandomToken() Token {
	return Token(random.String(64))
}
