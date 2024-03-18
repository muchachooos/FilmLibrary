package auth

import (
	"golang.org/x/crypto/bcrypt"
)

const cost = 10

func CompareHashPassword(hashedPass, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
}

func HashPassword(pass string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		return "", err
	}

	return string(hashedPass), nil
}
