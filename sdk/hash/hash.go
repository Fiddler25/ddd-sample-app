package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func Generate(rawString string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(rawString), bcrypt.DefaultCost)
	return string(hash)
}
