package util

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

// func HashPassword(password string) string{
// 	hashedPassword , _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	return string(hashedPassword)
// }

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16) // Salt size is 16 bytes
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// HashPassword hashes the password with a new salt, mimicking bcrypt behavior
func HashPassword(password string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}
	// Argon2 parameters: 1 iteration, 64MB memory, 4 parallel threads, and 32-byte output
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	// Return the salt and hash encoded together in a single string
	// The format is: base64(salt + hash)
	combined := append(salt, hash...)
	return base64.StdEncoding.EncodeToString(combined), nil
}

// VerifyPassword verifies the password against the stored hash
func VerifyPassword(storedHash string, password string) (bool, error) {
	// Decode the stored hash (which includes both salt and hash)
	decoded, err := base64.StdEncoding.DecodeString(storedHash)
	if err != nil {
		return false, err
	}

	// The first 16 bytes are the salt
	salt := decoded[:16]
	// The rest is the hash
	storedPasswordHash := decoded[16:]

	// Hash the input password with the same salt
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	// Compare the stored hash with the newly generated hash
	return compareHashes(storedPasswordHash, hash), nil
}

// Compare hashes securely (constant-time comparison)
func compareHashes(hash1, hash2 []byte) bool {
	if len(hash1) != len(hash2) {
		return false
	}
	for i := 0; i < len(hash1); i++ {
		if hash1[i] != hash2[i] {
			return false
		}
	}
	return true
}
