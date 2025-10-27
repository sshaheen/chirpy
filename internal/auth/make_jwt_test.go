package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeJWT(t *testing.T) {
	userId, err := uuid.NewUUID()

	if err != nil {
		t.Errorf("could not create user id")
	}

	tokenSecret := "mysamplesecret"
	_, err = MakeJWT(userId, tokenSecret, time.Duration(10))

	if err != nil {
		t.Errorf("error creating token")
	}
}
