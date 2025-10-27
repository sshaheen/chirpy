package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestValidateJWT(t *testing.T) {
	userId, err := uuid.NewUUID()

	if err != nil {
		t.Errorf("could not create user id")
	}

	tokenSecret := "mysamplesecret"
	tokenString, err := MakeJWT(userId, tokenSecret, time.Duration(time.Hour*1))

	if err != nil {
		t.Errorf("error creating token: %s", err)
	}

	_, err = ValidateJWT(tokenString, tokenSecret)
	if err != nil {
		t.Errorf("Token is invalid: %s", err)
	}
}

func TestValidateJWTWrongSecret(t *testing.T) {
	userId, err := uuid.NewUUID()

	if err != nil {
		t.Errorf("could not create user id")
	}

	tokenSecret := "mysamplesecret"
	tokenString, err := MakeJWT(userId, tokenSecret, time.Duration(time.Hour*1))

	if err != nil {
		t.Errorf("error creating token")
	}

	_, err = ValidateJWT(tokenString, "wrongsecret")
	if err == nil {
		t.Errorf("expected error with wrong secret")
	}
}

func TestValidateJWTExpired(t *testing.T) {
	userId, err := uuid.NewUUID()

	if err != nil {
		t.Errorf("could not create user id")
	}

	tokenSecret := "mysamplesecret"
	tokenString, err := MakeJWT(userId, tokenSecret, time.Duration(time.Microsecond*5))

	if err != nil {
		t.Errorf("error creating token")
	}

	_, err = ValidateJWT(tokenString, tokenSecret)
	if err == nil {
		t.Errorf("expected expired token error")
	}
}
