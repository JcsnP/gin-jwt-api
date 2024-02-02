package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/JcsnP/go-jwt-app/config/db"
	"github.com/JcsnP/go-jwt-app/config/schema"
	"github.com/golang-jwt/jwt"
)

func RemoveRecord() {
	db.DB.Where("1 = 1").Delete(&schema.User{})
}

func GenerateToken(iss uint) string {
	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": iss,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		panic(err.Error())
	}

	return tokenString
}

func ValidateToken(accessToken string) (interface{}, error) {
	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["iss"], nil
	}
	return nil, fmt.Errorf("invalid token")
}