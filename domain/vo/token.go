package vo

import (
	"github.com/Fiddler25/ddd-sample-app/sdk/hash"
	"github.com/labstack/gommon/random"
)

type Token string

func NewToken() Token {
	return Token(hash.Generate(random.String(64)))
}
