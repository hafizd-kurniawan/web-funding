package register

import (
	"funding/internal/identity/domain"

	"github.com/google/uuid"
)

type Handler struct {
	repo   domain.UserRepository
	hasher domain.Hasher
}

func NewHandler(repo domain.UserRepository, hasher domain.Hasher) *Handler {
	return &Handler{
		repo:   repo,
		hasher: hasher,
	}
}

func (h *Handler) Handle(cmd Command) (Result, error) {
	email, err := domain.NewEmail(cmd.Email)
	if err != nil {
		return Result{}, err
	}

	_, err = h.repo.FindByEmail(cmd.Email)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return Result{}, err
	}

	password, err := domain.NewPassword(cmd.Password, h.hasher)
	if err != nil {
		return Result{}, err
	}

	userID := domain.UserID(uuid.New())
	user, _, err := domain.RegisterUser(userID, cmd.Name, email.Value())
	if err != nil {
		return Result{}, err
	}
	user.ChangePassword(password)

	if err := h.repo.Save(user); err != nil {
		return Result{}, err
	}
	return Result{Name: user.GetName(), Email: user.GetEmail()}, nil
}
