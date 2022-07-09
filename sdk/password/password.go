package password

import (
	"golang.org/x/crypto/bcrypt"
)

func IsSame(currPassword, reqPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(currPassword), []byte(reqPassword))
	return err == nil
}
