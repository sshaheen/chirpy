package auth

import (
	"testing"

	"github.com/alexedwards/argon2id"
)

func TestCheckPasswordHash(t *testing.T) {
	password := "password"
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		t.Errorf("Failed to create hash.\n")
	}
	_, err = CheckPasswordHash(password, hash)
	if err != nil {
		t.Errorf("CheckPasswordHash failed.\n")
	}
}
