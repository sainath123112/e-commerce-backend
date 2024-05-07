package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string) (string, error) {
	secret := []byte("secret string")

	tokenCliams := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"iss":      time.Now(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := tokenCliams.SignedString(secret)
	return token, err
}

func ValidateToken(JwtToken string) (bool, error) {
	token, err := ParseToken(JwtToken)
	if err != nil {
		return false, err
	}
	if !token.Valid {
		return false, jwt.ErrInvalidKey
	}
	return true, nil
}

func ParseToken(JwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(JwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret string"), nil
	})

	return token, err
}

func GetUsername(JwtToken string) (string, error) {
	token, err := ParseToken(JwtToken)
	if err != nil {
		return "Unable to parse token", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "Unable to get Claims", jwt.ErrTokenInvalidClaims
	}
	username := claims["username"].(string)
	return username, nil
}
