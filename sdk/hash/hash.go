package hash

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Generate(rawString string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawString), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}
