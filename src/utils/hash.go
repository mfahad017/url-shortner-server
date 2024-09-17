package utils

import "golang.org/x/crypto/bcrypt"

type Password string

// HashPassword hashes the given password using bcrypt and returns the hashed password.
func (p Password) HashPassword() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compares the given password with the hashed password using bcrypt and returns true if they match.
func (p Password) ComparePassword(hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(p))
	return err == nil
}
