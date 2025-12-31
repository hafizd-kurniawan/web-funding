package user

import (
	"errors"
	"time"
)

type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          Email
	Password       Password
	AvatarFileName string
	Role           Role
	IsActive       bool
	CreatedAt      time.Time
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
	IsDeleted      bool

	events []any
}

func RegisterUser(name, email, password string) (*User, error) {
	if name == "" || len(name) < 3 {
		return nil, ErrNameNotValid
	}

	e, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	p, err := NewPasswordFromPlain(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:             1,
		Name:           name,
		Occupation:     "",
		Email:          e,
		Password:       p,
		AvatarFileName: "",
		Role:           RoleUser,
		IsActive:       true,
		CreatedAt:      time.Now(),
		IsDeleted:      false,
	}

	user.events = append(user.events, UserRegistered{
		Email: user.Email.String(),
		Name:  user.Name,
		At:    time.Now(),
	})
	return user, nil
}

func (u *User) ChangePassword(oldPassword, newPassword string) error {
	if !u.Password.Match(oldPassword) {
		return ErrPasswordDoesNotMatch
	}

	pass, err := NewPasswordFromPlain(newPassword)
	if err != nil {
		return err
	}

	u.Password = pass
	return nil
}

func (u *User) IsUserActive() bool {
	return u.IsActive
}

func (u *User) UpdateAvatar(fileName string) error {
	if !u.IsUserActive() {
		return errors.New("user is not active")
	}

	if fileName == "" {
		return errors.New("fileName empty")
	}

	u.AvatarFileName = fileName
	return nil
}

func (u *User) PullEvents() []any {
	ev := u.events
	u.events = nil
	return ev
}
