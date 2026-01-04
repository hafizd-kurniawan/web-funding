package application

import (
	"errors"
	"funding/internal/identity/domain"
	"time"

	"github.com/google/uuid"
)

type RegisterUserCommand struct {
	Name     string
	Email    string
	Password string
}

type LoginUserCommand struct {
	Email    string
	Password string
}

type RefreshTokenCommand struct {
	RefreshToken uuid.UUID
}

type LoginUserResult struct {
	Name         string
	Occupation   string
	Email        string
	Token        string
	RefreshToken uuid.UUID
}

type ChangePasswordCommand struct {
	Email    string
	Password string
}

type ChangeAvatarCommand struct {
	Email  string
	Avatar string
}

type UpdateUserCommand struct {
	Name       string
	Email      string
	Occupation string
}

type RegisterUser struct {
	repo         domain.UserRepository
	hasher       domain.Hasher
	token        domain.TokenGenerator
	refreshToken domain.RefreshTokenRepo
}

func (uc *RegisterUser) RegisterUser(cmd RegisterUserCommand) error {
	_, err := uc.repo.FindByEmail(cmd.Email)
	if err == nil {
		return errors.New("email already exist")
	}

	email, err := domain.NewEmail(cmd.Email)
	if err != nil {
		return err
	}

	userID := domain.UserID(uuid.New())
	user, _, err := domain.RegisterUser(userID, cmd.Name, email.Value())
	if err != nil {
		return err
	}

	password, err := domain.NewPassword(cmd.Password, uc.hasher)
	if err != nil {
		return err
	}
	user.ChangePassword(password)

	return uc.repo.Save(user)
}

func (uc *RegisterUser) LoginUser(cmd RegisterUserCommand) (LoginUserResult, error) {
	var result LoginUserResult
	u, err := uc.repo.FindByEmail(cmd.Email)
	if err != nil {
		return result, errors.New("user not found")
	}

	isValid := uc.hasher.Compare(cmd.Password, u.GetPassword().Value())
	if !isValid {
		return result, errors.New("password is invalid")
	}

	token, err := uc.token.GenerateToken(u.GetID())
	if err != nil {
		return result, err
	}

	refreshToken := domain.RefreshToken{
		ID:        uuid.New(),
		UserID:    u.GetID(),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
		Revoked:   false,
	}

	if err := uc.refreshToken.Save(refreshToken); err != nil {
		return result, err
	}

	result = LoginUserResult{
		Name:         u.GetName(),
		Occupation:   u.GetOccupation(),
		Email:        u.GetEmail(),
		Token:        token,
		RefreshToken: refreshToken.ID,
	}
	return result, nil
}

func (uc *RegisterUser) RefreshToken(cmd RefreshTokenCommand) (string, string, error) {
	rt, err := uc.refreshToken.FindByID(cmd.RefreshToken)
	if err != nil {
		return "", "", err
	}

	if rt.Revoked || rt.IsExpired() {
		return "", "", errors.New("refresh token is revoked or expired")
	}

	if err = uc.refreshToken.Revoke(rt.ID); err != nil {
		return "", "", err
	}

	newRT := domain.RefreshToken{
		ID:        uuid.New(),
		UserID:    rt.UserID,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
		Revoked:   false,
	}

	if err = uc.refreshToken.Save(newRT); err != nil {
		return "", "", err
	}

	accessToken, err := uc.token.GenerateToken(rt.UserID)
	if err != nil {
		return "", "", err
	}

	return accessToken, newRT.ID.String(), nil
}

func (uc *RegisterUser) ChangePassword(cmd RegisterUserCommand) error {
	u, err := uc.repo.FindByEmail(cmd.Email)
	if err != nil {
		return errors.New("user not found")
	}

	password, err := domain.NewPassword(cmd.Password, uc.hasher)
	if err != nil {
		return err
	}
	if err := u.ChangePassword(password); err != nil {
		return err
	}

	return uc.repo.Save(u)
}

func (uc *RegisterUser) ChangeAvatar(cmd ChangeAvatarCommand) error {
	u, err := uc.repo.FindByEmail(cmd.Email)
	if err != nil {
		return errors.New("email not found")
	}

	u.ChangeAvatar(cmd.Avatar)
	return uc.repo.Save(u)
}

func (uc *RegisterUser) UpdateUser(cmd UpdateUserCommand) error {
	u, err := uc.repo.FindByEmail(cmd.Email)
	if err != nil {
		return errors.New("email already exist")
	}
	if err := u.UpdateUser(cmd.Name, cmd.Email, cmd.Occupation); err != nil {
		return err
	}
	return uc.repo.Save(u)
}
