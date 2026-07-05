package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"strings"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {
	// Generate a random 16-byte salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,             
		1,                
		64*1024,          
		4,                
		32,               
	)

	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	return encodedSalt + "." + encodedHash, nil
}

func VerifyPassword(input string, comparedPassword string) bool {
	parts := strings.Split(comparedPassword, ".")
	if len(parts) != 2 {
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false
	}

	storedHash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false
	}

	hash := argon2.IDKey(
		[]byte(input),
		salt,
		1,
		64*1024,
		4,
		uint32(len(storedHash)),
	)

	return subtle.ConstantTimeCompare(hash, storedHash) == 1
}