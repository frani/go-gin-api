package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

const (
	saltSize   = 16    // Size of the salt in bytes
	keyLength  = 32    // Length of the derived key in bytes
	iterations = 10000 // Number of iterations for PBKDF2
)

func HashPassword(password string) ([]byte, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	hashedPassword := pbkdf2.Key([]byte(password), salt, iterations, keyLength, sha256.New)
	hashedPasswordWithSalt := append(salt, hashedPassword...)
	return hashedPasswordWithSalt, nil
}

func VerifyPassword(password string, hashedPassword []byte) bool {
	salt := hashedPassword[:saltSize]
	computedHash := pbkdf2.Key([]byte(password), salt, iterations, keyLength, sha256.New)
	return hmac.Equal(hashedPassword[saltSize:], computedHash)
}
