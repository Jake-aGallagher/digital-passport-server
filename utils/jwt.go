package utils

import (
	"errors"
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

func VerifyToken(token string) (string, string, error) {
	var err = godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
	secretKey := os.Getenv("SECRETKEY")
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", "", errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return "", "", errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("invalid token claims")
	}

	userId := claims["userId"].(string)
	companyId := claims["companyId"].(string)

	return userId, companyId, nil
}
