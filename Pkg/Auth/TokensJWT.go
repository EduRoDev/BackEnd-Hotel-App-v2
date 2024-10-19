package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret")

func GenerateToken(userId int, userName string) (string, error) {
	claims := jwt.MapClaims{
		"id":   userId,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"name": userName,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

