package pwdcrypto

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type ArgonParams struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var defaultParams = ArgonParams{
	Memory:      64 * 1024, // 64 MB
	Iterations:  1,
	Parallelism: 1,
	SaltLength:  16,
	KeyLength:   32,
}

func HashPassword(password string) (string, error) {
	salt := make([]byte, defaultParams.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		defaultParams.Iterations,
		defaultParams.Memory,
		defaultParams.Parallelism,
		defaultParams.KeyLength,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		defaultParams.Memory,
		defaultParams.Iterations,
		defaultParams.Parallelism,
		b64Salt,
		b64Hash,
	)

	return encodedHash, nil
}
