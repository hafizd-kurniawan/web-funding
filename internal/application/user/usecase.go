package user

import (
	"context"
	"errors"
	"funding/internal/application/event"
	"funding/internal/domain/user"
	"funding/internal/infrastructure/persistence/repository"
)

type UserUseCase struct {
	repo       user.Repository
	events     event.Publisher
	repoOutbox repository.WorkerRepository
}

func NewRegisterUserUseCase(r user.Repository, events event.Publisher) *UserUseCase {
	return &UserUseCase{repo: r, events: events}
}

func (uc *UserUseCase) Register(ctx context.Context, cmd RegisterUserCommand) error {

	email, err := user.NewEmail(cmd.Email)
	if err != nil {
		return err
	}

	_, err = uc.repo.FindByEmail(ctx, email)
	if err == nil {
		return errors.New("email already exist")
	}

	user, err := user.RegisterUser(
		cmd.Name,
		cmd.Email,
		cmd.Password,
	)

	if err != nil {
		return err
	}

	if err := uc.repo.Save(ctx, user); err != nil {
		return err
	}

	// publish domain events
	for _, ev := range user.PullEvents() {
		uc.events.Publish(ctx, ev)
	}

	return nil
}

func (uc *UserUseCase) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	return uc.repo.FindByEmail(ctx, user.Email(email))
}

func (uc *UserUseCase) FindByID(ctx context.Context, id int) (*user.User, error) {
	return uc.repo.FindByID(ctx, id)
}
