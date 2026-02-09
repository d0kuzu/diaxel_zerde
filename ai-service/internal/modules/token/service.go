package token

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateSecureToken(prefix string, byteLength int) (string, error) {
	b := make([]byte, byteLength)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	token := base64.RawURLEncoding.EncodeToString(b)

	return fmt.Sprintf("%s_%s", prefix, token), nil
}
