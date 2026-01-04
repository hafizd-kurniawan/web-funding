package domain

import (
	"errors"
	"strings"
)

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	e := Email{}
	if value == "" {
		return e, errors.New("email is required")
	}

	if !strings.Contains(value, "@") {
		return e, errors.New("invalid email format")
	}

	e.value = value
	return e, nil
}

func (e Email) Value() string {
	return e.value
}
