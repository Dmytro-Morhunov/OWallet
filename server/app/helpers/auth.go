package helpers

import (
	"crypto/rand"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var SecretKey = make([]byte, 64)

func CreateToken(email string) (string, error) {
	_, generateError := rand.Read(SecretKey)
	if generateError != nil {
		return "", generateError
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	token, err := claims.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Invalid JWT Token")
}
