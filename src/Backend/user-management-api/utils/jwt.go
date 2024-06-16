package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "FGNMfvbnmVBNM23456789$%^&*(*&^%$%^&*&^%$%^&^%$%^&^%$#)"

func GenerateToken(username string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"userId":   userId,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
