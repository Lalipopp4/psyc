package scripts

import (
	"os"
	"psyc/internal/errors"
	"strings"

	"github.com/golang-jwt/jwt"
)

var (
	// Very secret key
	key = []byte(os.Getenv("JWT_SECRET"))
)

type claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// Generates JWT for user
func GenerateJWTUserToken(id, email string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    id,
			"email": email,
		},
	).SignedString(key)
}

// Parses user token
func ParseJWTUserToken(accessToken string) (string, string, error) {
	accessToken = strings.Split(accessToken, " ")[1]
	token, err := jwt.ParseWithClaims(
		accessToken,
		&claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.ErrorServer{}
			}
			return key, nil
		})
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(*claims)
	if !ok {
		return "", "", err
	}
	return claims.ID, claims.Email, nil
}
