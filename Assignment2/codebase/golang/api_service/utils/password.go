package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	defaultHashCost = bcrypt.DefaultCost
)

// HashPassword generates a hashed version of the password
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), defaultHashCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPassword checks if the password is correct
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
