package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateToken(userId string, companyId string) (string, error) {
	var err = godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
	secretKey := os.Getenv("SECRETKEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"companyId": companyId,
		"userId":    userId,
		"exp":       time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
