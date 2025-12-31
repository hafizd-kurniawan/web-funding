package user

import (
	"net/mail"
)

type Email string

func NewEmail(v string) (Email, error) {
	e := Email(v)

	if err := e.validate(); err != nil {
		return "", err
	}

	return e, nil
}

func (e Email) validate() error {
	if e == "" {
		return ErrEmailNotValid
	}

	_, err := mail.ParseAddress(string(e))
	return err
}

func (e Email) String() string {
	return string(e)
}
