package user_test

import (
	"context"
	"errors"
	"testing"


	"funding/internal/application/user"
	domain "funding/internal/domain/user"
)

// MockRepository is a mock implementation of user.Repository
type MockRepository struct {
	users map[string]*domain.User
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		users: make(map[string]*domain.User),
	}
}

func (m *MockRepository) Save(ctx context.Context, u *domain.User) error {
	m.users[u.Email.String()] = u
	return nil
}

func (m *MockRepository) FindByEmail(ctx context.Context, email domain.Email) (*domain.User, error) {
	if u, ok := m.users[email.String()]; ok {
		return u, nil
	}
	return nil, errors.New("user not found")
}

func (m *MockRepository) FindByID(ctx context.Context, id int) (*domain.User, error) {
	return nil, nil
}

// MockPublisher is a mock implementation of event.Publisher
type MockPublisher struct {
	PublishedEvents []any
}

func (m *MockPublisher) Publish(ctx context.Context, event any) error {
	m.PublishedEvents = append(m.PublishedEvents, event)
	return nil
}

func TestRegisterUser(t *testing.T) {
	repo := NewMockRepository()
	publisher := &MockPublisher{}
	useCase := user.NewRegisterUserUseCase(repo, publisher)

	cmd := user.RegisterUserCommand{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	err := useCase.Register(context.Background(), cmd)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify user was saved
	savedUser, err := repo.FindByEmail(context.Background(), domain.Email("john@example.com"))
	if err != nil {
		t.Fatalf("expected user to be saved, got error %v", err)
	}
	if savedUser.Name != "John Doe" {
		t.Errorf("expected name %s, got %s", "John Doe", savedUser.Name)
	}

	// Verify event was published
	if len(publisher.PublishedEvents) != 1 {
		t.Fatalf("expected 1 event, got %d", len(publisher.PublishedEvents))
	}

	event, ok := publisher.PublishedEvents[0].(domain.UserRegistered)
	if !ok {
		t.Fatalf("expected event of type UserRegistered, got %T", publisher.PublishedEvents[0])
	}

	if event.Email != "john@example.com" {
		t.Errorf("expected event email %s, got %s", "john@example.com", event.Email)
	}
}
