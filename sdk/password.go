package sdk

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(rawPassword string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}
