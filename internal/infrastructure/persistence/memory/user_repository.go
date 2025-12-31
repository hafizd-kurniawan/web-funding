package memory

import (
	"context"
	"errors"
	"funding/internal/domain/user"
	"sync"
)

type InMemoryUserRepository struct {
	mu    sync.Mutex
	users map[string]*user.User
}

var _ user.Repository = (*InMemoryUserRepository)(nil)

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*user.User),
	}
}

func (r *InMemoryUserRepository) Save(ctx context.Context, u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[u.Email.String()] = u
	return nil
}

func (r *InMemoryUserRepository) FindByEmail(ctx context.Context, email user.Email) (*user.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if u, ok := r.users[email.String()]; ok {
		return u, nil
	}
	return nil, errors.New("user not found")
}

func (r *InMemoryUserRepository) FindByID(ctx context.Context, id int) (*user.User, error) {
	return nil, nil
}
