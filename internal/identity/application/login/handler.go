package login

import (
	"errors"
	"fmt"
	"funding/internal/identity/domain"
	"time"

	"github.com/google/uuid"
)

type Handler struct {
	repo         domain.UserRepository
	hasher       domain.Hasher
	accessToken  domain.TokenGenerator
	refreshToken domain.RefreshTokenRepo
}

func NewHandler(repo domain.UserRepository, hasher domain.Hasher, accessToken domain.TokenGenerator, refreshToken domain.RefreshTokenRepo) *Handler {
	return &Handler{
		repo:         repo,
		hasher:       hasher,
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}
}

func (h *Handler) Handle(cmd Command) (Result, error) {
	user, err := h.repo.FindByEmail(cmd.Email)
	if err != nil {
		return Result{}, err
	}

	isValid := h.hasher.Compare(cmd.Password, user.GetPassword().Value())
	if !isValid {
		return Result{}, errors.New("password is invalid")
	}

	token, err := h.accessToken.GenerateToken(user.GetID())

	if err != nil {
		return Result{}, err
	}

	fmt.Printf("%+v\n", user)

	refreshToken := domain.RefreshToken{
		ID:        uuid.New(),
		UserID:    user.GetID(),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
		Revoked:   false,
	}

	if err := h.refreshToken.Save(refreshToken); err != nil {
		return Result{}, err
	}

	result := Result{
		Name:         user.GetName(),
		Email:        user.GetEmail(),
		Token:        token,
		RefreshToken: refreshToken.ID,
	}
	return result, nil
}
