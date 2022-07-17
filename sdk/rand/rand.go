package rand

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(n int) (string, error) {
	if n < 0 {
		return "", nil
	}
	b, err := generateRandomBytes(n)

	return base64.URLEncoding.EncodeToString(b), err
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)

	return b, err
}
