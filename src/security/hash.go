package security

import (
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func GeneratePasswordHash(password string) string {
	hash := argon2.IDKey([]byte(password), []byte(Secret), 1, 64*512, 2, 32)
	encodedHash := base64.URLEncoding.EncodeToString(hash)
	return encodedHash
}
