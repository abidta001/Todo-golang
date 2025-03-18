package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecretKey = "This is my secretkey"

func GenerateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 5).Unix(),
	})
	return token.SignedString([]byte(JWTSecretKey))
}
