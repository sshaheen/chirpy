package auth

import (
	"crypto/rand"
	"encoding/hex"
)

func MakeRefreshToken() (string, error) {
	tokStr := make([]byte, 32)
	_, err := rand.Read(tokStr)
	if err != nil {
		return "", err
	}

	encodedHex := hex.EncodeToString(tokStr)
	return encodedHex, nil
}
