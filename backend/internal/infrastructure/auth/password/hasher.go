package password

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

// Hasher represents password hasher using Argon2
type Hasher struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

// NewHasher creates a new password hasher
func NewHasher() *Hasher {
	return &Hasher{
		memory:      64 * 1024, // 64 MB
		iterations:  1,
		parallelism: 4,
		saltLength:  24,
		keyLength:   32,
	}
}

// HashPassword hashes a password using Argon2
func (h *Hasher) HashPassword(password string) (string, string, error) {
	salt := make([]byte, h.saltLength)
	rand.Read(salt)

	hash := argon2.IDKey([]byte(password), salt, h.iterations, h.memory, h.parallelism, h.keyLength)
	encodedPassword := base64.RawStdEncoding.EncodeToString(hash)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)

	return encodedPassword, encodedSalt, nil
}

// VerifyPassword verifies a password against its hash
func (h *Hasher) VerifyPassword(password, encodedPassword, salt string) (bool, error) {
	decodedSalt, err := base64.RawStdEncoding.Strict().DecodeString(salt)
	if err != nil {
		return false, err
	}

	decodedPassword, err := base64.RawStdEncoding.Strict().DecodeString(encodedPassword)
	if err != nil {
		return false, err
	}

	hash := argon2.IDKey([]byte(password), decodedSalt, h.iterations, h.memory, h.parallelism, h.keyLength)
	return subtle.ConstantTimeCompare(decodedPassword, hash) == 1, nil
}
