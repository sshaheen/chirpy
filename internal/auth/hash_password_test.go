package auth

import (
	"testing"

	"github.com/alexedwards/argon2id"
)

func TestHashPassword(t *testing.T) {
	password := "password"
	output, hashErr := HashPassword(password)
	result, err := argon2id.ComparePasswordAndHash(password, output)
	if err != nil || hashErr != nil {
		t.Errorf("HashPassword return %v.\n", result)
	}
}
