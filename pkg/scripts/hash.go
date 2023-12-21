package scripts

import (
	crypto "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := crypto.GenerateFromPassword([]byte(password), 10)
	return string(hash), err
}

func CheckPasswordHash(password, hashed string) bool {
	err := crypto.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
