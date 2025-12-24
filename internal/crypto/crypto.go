package _crypto

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashFromPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, fmt.Errorf("an error occurred in password generation: %w", err)
	}

	return hash, nil
}
