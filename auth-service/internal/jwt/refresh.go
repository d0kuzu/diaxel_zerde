package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateRefreshToken(userID string, ttl time.Duration, secret string) (string, time.Time, error) {
	expirationTime := time.Now().Add(ttl)
	claims := jwt.MapClaims{
		"sub":  userID,
		"type": "refresh",
		"exp":  expirationTime.Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(secret))
	return token, expirationTime, err
}

func ParseRefreshToken(token, secret string) (string, error) {
	parsed, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil || !parsed.Valid {
		return "", err
	}

	claims := parsed.Claims.(jwt.MapClaims)
	return claims["sub"].(string), nil
}
