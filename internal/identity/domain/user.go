package domain

import (
	"database/sql/driver"
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
	id         UserID
	name       string
	email      Email
	password   Password
	role       Role
	occupation string
	avatar     string
	isActive   bool
	isVerified bool
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

	u.id = id
	u.name = name
	u.email = e
	u.occupation = "Developer"
	u.isActive = true
	u.isVerified = true
	u.role = RoleUser

	return u, []any{UserRegistered{ID: u.id, Email: u.email.Value()}}, nil
}

func (u *User) ChangePassword(password Password) error {
	if !u.IsUserActive() {
		return errors.New("user is not active")
	}

	if len(password.value) == 0 {
		return errors.New("password is required")
	}

	u.password = password
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

	u.name = name
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

	u.occupation = occupation
	return nil
}

func (u *User) DeactivateUser() error {
	if !u.IsUserActive() {
		return errors.New("user is not active")
	}

	u.isActive = false
	return nil
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

	u.name = name
	u.email = e
	u.occupation = occupation
	return nil
}

func (u *User) IsUserActive() bool    { return u.isActive }
func (u *User) IsUserVerified() bool  { return u.isVerified }
func (u *User) GetAvatar() string     { return u.avatar }
func (u *User) GetRole() Role         { return u.role }
func (u *User) GetOccupation() string { return u.occupation }
func (u *User) GetName() string       { return u.name }
func (u *User) GetID() UserID         { return u.id }
func (u *User) GetUUID() uuid.UUID    { return uuid.UUID(u.id) }
func (u *User) GetEmail() string      { return u.email.Value() }
func (u *User) GetPassword() Password { return u.password }
func (id UserID) String() string {
	uuid := uuid.UUID(id)
	return uuid.String()
}

// Scan implements the sql.Scanner interface
func (id *UserID) Scan(value interface{}) error {
	var u uuid.UUID
	if err := u.Scan(value); err != nil {
		return err
	}
	*id = UserID(u)
	return nil
}

// Value implements the driver.Valuer interface
func (id UserID) Value() (driver.Value, error) {
	return uuid.UUID(id).Value()
}
