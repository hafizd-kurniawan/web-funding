package hasher

import (
	"funding/internal/identity/domain"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct{}

var _ domain.Hasher = BcryptHasher{}

func (h BcryptHasher) Hash(plain string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
}

func (h BcryptHasher) Compare(plain, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}
