package auth

import (
	"github.com/Fiddler25/sample-app/sdk/hash"
	"golang.org/x/crypto/bcrypt"
)

type Password string

func NewPassword(rawValue string) Password {
	return Password(hash.Generate(rawValue))
}

func IsSamePassword(currPassword, reqPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(currPassword), []byte(reqPassword))
	return err == nil
}
