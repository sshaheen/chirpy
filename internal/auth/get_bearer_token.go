package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("no authorization header")
	}
	authHeaderSlice := strings.Split(authHeader, " ")
	if len(authHeaderSlice) < 2 {
		return "", fmt.Errorf("auth header could not split")
	}
	tokenString := strings.TrimSpace(authHeaderSlice[1])
	return tokenString, nil
}
