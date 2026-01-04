package domain

import "errors"

type Hasher interface {
	Hash(plain string) ([]byte, error)
	Compare(plain, hashed string) bool
}

type Password struct {
	value string
}

func NewPassword(value string, hasher Hasher) (Password, error) {
	p := Password{}

	if value == "" {
		return p, errors.New("password is required")
	}

	if len(value) < 6 {
		return p, errors.New("password must be at least 8 characters")
	}

	hashed, err := hasher.Hash(value)
	if err != nil {
		return p, err
	}

	p.value = string(hashed)
	return p, nil
}

func (p Password) Value() string {
	return p.value
}
