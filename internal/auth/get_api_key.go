package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authStr := headers.Get("Authorization")
	splitAuthStr := strings.Split(authStr, " ")
	if len(splitAuthStr) < 2 {
		return "", fmt.Errorf("auth header could not split")
	}
	apiKey := strings.TrimSpace(splitAuthStr[1])
	return apiKey, nil
}
