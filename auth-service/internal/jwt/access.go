package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userID, role string, ttl time.Duration, secret string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"type": "access",
		"exp":  time.Now().Add(ttl).Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(secret))
}
