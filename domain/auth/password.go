package auth

import (
	"github.com/Fiddler25/ddd-sample-app/sdk/hash"
)

type Password string

func NewPassword(rawValue string) Password {
	return Password(hash.Generate(rawValue))
}
