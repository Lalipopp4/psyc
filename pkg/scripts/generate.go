package scripts

import "math/rand"

const (
	letters = "abcdefghijklmnopqrstuvwxyABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// Generates random id
func GenerateID() string {
	id := make([]byte, 16)
	for i := range id {
		id[i] = letters[rand.Intn(61)]
	}
	return string(id)
}
