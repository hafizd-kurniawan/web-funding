package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	hash string
}

func NewPasswordFromPlain(raw string) (Password, error) {
	if len(raw) < 8 {
		return Password{}, errors.New("password too short")
	}

	hashed, err := createHashed(raw)
	if err != nil {
		return Password{}, err
	}

	return Password{hash: string(hashed)}, nil
}

func NewPasswordFromHash(hash string) Password {
	return Password{hash: hash}
}

func createHashed(raw string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return hashed, err
	}
	return hashed, nil
}

func (p Password) Match(raw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.hash), []byte(raw))
	return err == nil
}

func (p Password) Hash() string {
	return p.hash
}
