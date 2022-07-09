package vo

import (
	"github.com/Fiddler25/ddd-sample-app/sdk/password"
)

type Password string

func NewPassword(rawValue string) Password {
	return Password(password.Hash(rawValue))
}
