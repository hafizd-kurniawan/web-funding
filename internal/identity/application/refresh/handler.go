package refresh

import (
	"errors"
	"fmt"
	"funding/internal/identity/domain"
)

type Handler struct {
	refreshToken domain.RefreshTokenRepo
	accessToken  domain.TokenGenerator
}

func NewHandler(refreshToken domain.RefreshTokenRepo, accessToken domain.TokenGenerator) *Handler {
	return &Handler{
		refreshToken: refreshToken,
		accessToken:  accessToken,
	}
}

func (h *Handler) Handle(cmd Command) (Result, error) {
	rt, err := h.refreshToken.FindByID(cmd.UserID)
	if err != nil {
		fmt.Println("error disini")
		return Result{}, err
	}
	if rt.Revoked {
		return Result{}, errors.New("refresh token is revoked")
	}
	if err = h.refreshToken.Revoke(rt.ID); err != nil {
		return Result{}, err
	}

	newRT := domain.NewRefreshTokenID(rt.UserID)
	if err = h.refreshToken.Save(newRT); err != nil {
		return Result{}, err
	}

	accessToken, err := h.accessToken.GenerateToken(rt.UserID)
	if err != nil {
		return Result{}, err
	}

	return Result{
		AccessToken:  accessToken,
		RefreshToken: newRT.ID.String(),
	}, nil
}
