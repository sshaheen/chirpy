package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claims := &jwt.RegisteredClaims{}
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)

	if err != nil {
		return uuid.UUID{}, err
	}

	idStr, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.UUID{}, err
	}

	resUUID, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.UUID{}, err
	}

	return resUUID, nil
}
