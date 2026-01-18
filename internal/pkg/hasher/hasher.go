package hasher

import (
	"errors"
	"strconv"

	"github.com/AurChatOrg/aurchat-server/internal/code"
	"github.com/AurChatOrg/aurchat-server/internal/pkg/logger"
	"github.com/alexedwards/argon2id"
	"go.uber.org/zap"
)

type Hasher struct {
	argon2 *argon2id.Params
}

// NewHasher New argon2 hash instance
func NewHasher(memory, iterations uint32, parallelism uint8, saltLength, keyLength uint32) *Hasher {
	return &Hasher{
		argon2: &argon2id.Params{
			Memory:      memory,
			Iterations:  iterations,
			Parallelism: parallelism,
			SaltLength:  saltLength,
			KeyLength:   keyLength,
		},
	}
}

// Hash Convert strings to hashes using Argon2
func (a *Hasher) Hash(plaintext string) string {
	hash, err := argon2id.CreateHash(plaintext, a.argon2) // Generate Argon2 hash

	if err != nil {
		logger.Logger.Fatal("Create hash error", zap.Error(err))
	}
	return hash
}

// VerifyHash Verify if the string matches the provided hash
func (a *Hasher) VerifyHash(plaintext string, hash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(plaintext, hash) // Verify hash
	if err != nil {
		logger.Logger.Fatal("Compare hash error", zap.Error(err))
		return false, errors.New(strconv.Itoa(code.ErrorAccountNameOrPassword))
	}

	return match, nil
}
