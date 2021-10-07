package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	hashedPassword := string(hash)

	return hashedPassword
}

func CheckPassword(userPassword string, correctPassword []byte) bool {
	if err := bcrypt.CompareHashAndPassword(correctPassword, []byte(userPassword)); err != nil {
		return false
	}

	return true
}
