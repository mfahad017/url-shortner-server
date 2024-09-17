package utils

import (
	"server/src/config"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JWT struct{}

func (JWT) GenerateToken(userId uint) (string, error) {
	// Define the secret key used for signing the token

	// Create a new token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Add(time.Hour * time.Duration(config.GetConfig().AUTH_TOKEN_EXPIRY_IN_MINUTES)).Unix(),
		"userId": userId,
	})

	// Generate the token string
	tokenString, err := token.SignedString([]byte(config.GetConfig().JWT_SECRET))
	if err != nil {
		// Handle error if token signing fails
		return "", err
	}

	return tokenString, nil
}
func (JWT) VerifyToken() bool {
	return true
}
