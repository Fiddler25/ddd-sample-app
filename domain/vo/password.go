package vo

import "github.com/Fiddler25/ddd-sample-app/sdk"

type Password string

func NewPassword(rawValue string) Password {
	return Password(sdk.HashPassword(rawValue))
}
