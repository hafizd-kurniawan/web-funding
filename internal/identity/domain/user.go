package domain

import (
	"errors"

	"github.com/google/uuid"
)

type (
	UserID uuid.UUID
	Role   string
)

const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
)

type User struct {
	ID         UserID
	Name       string
	Email      Email
	Password   Password
	Role       Role
	Occupation string
	avatar     string
	IsActive   bool
	IsVerified bool
}

func RegisterUser(id UserID, name, email string) (*User, []any, error) {
	u := new(User)

	if name == "" {
		return u, nil, errors.New("name is required")
	}

	if len(name) < 3 {
		return u, nil, errors.New("name must be at least 3 characters")
	}

	e, err := NewEmail(email)
	if err != nil {
		return u, nil, err
	}

	u.ID = id
	u.Name = name
	u.Email = e
	u.Occupation = "Developer"
	u.IsActive = true
	u.IsVerified = true
	u.Role = RoleUser

	return u, []any{UserRegistered{ID: u.ID, Email: u.Email.Value()}}, nil
}

func (u *User) ChangePassword(password Password) error {
	if !u.IsUserActive() {
		return errors.New("user is not active")
	}

	if len(password.value) == 0 {
		return errors.New("password is required")
	}

	u.Password = password
	return nil
}

func (u *User) ChangeAvatar(avatar string) error {
	if !u.IsUserActive() {
		return errors.New("user is not active")
	}

	if avatar == "" {
		return errors.New("avatar is required")
	}

	if len(avatar) < 3 {
		return errors.New("avatar must be at least 3 characters")
	}

	u.avatar = avatar
	return nil
}

func (u *User) ChangeName(name string) error {
	if !u.IsUserActive() {
		return errors.New("user is not active")
	}

	if name == "" {
		return errors.New("name is required")
	}

	if len(name) < 3 {
		return errors.New("name must be at least 3 characters")
	}

	if len(name) > 100 {
		return errors.New("name must be less than 100 characters")
	}

	u.Name = name
	return nil
}

func (u *User) ChangeOccupation(occupation string) error {
	if !u.IsUserActive() {
		return errors.New("user is not active")
	}

	if occupation == "" {
		return errors.New("occupation is required")
	}

	if len(occupation) < 3 {
		return errors.New("occupation must be at least 3 characters")
	}

	if len(occupation) > 100 {
		return errors.New("occupation must be less than 100 characters")
	}

	u.Occupation = occupation
	return nil
}

func (u *User) DeactivateUser() error {
	if !u.IsUserActive() {
		return errors.New("user is not active")
	}

	u.IsActive = false
	return nil
}

func (u *User) IsUserActive() bool {
	return u.IsActive
}

func (u *User) IsUserVerified() bool {
	return u.IsVerified
}

func (u *User) UpdateUser(name, email, occupation string) error {
	if !u.IsUserActive() {
		return errors.New("user is not active")
	}

	if name == "" {
		return errors.New("name is required")
	}

	if len(name) < 3 {
		return errors.New("name must be at least 3 characters")
	}

	e, err := NewEmail(email)
	if err != nil {
		return err
	}

	u.Name = name
	u.Email = e
	u.Occupation = occupation
	return nil
}

func (u *User) GetOccupation() string { return u.Occupation }
func (u *User) GetName() string       { return u.Name }
func (u *User) GetID() UserID         { return u.ID }
func (u *User) GetEmail() string      { return u.Email.Value() }
func (u *User) GetPassword() Password { return u.Password }
