package register

import (
	"testing"

	"funding/internal/identity/infrastructure/hasher"
	"funding/internal/identity/infrastructure/persistence/postgres"
	conn "funding/shared/database/postgres"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	db := conn.SetupDB()
	repo := postgres.NewRepository(db)
	hasher := hasher.BcryptHasher{}

	handler := NewHandler(repo, hasher)

	result, err := handler.Handle(Command{
		Name:     "Ayo",
		Email:    "ayo@gmail.commmmm",
		Password: "password123",
	})
	assert.NoError(t, err)
	assert.Equal(t, "ayo@gmail.commmmm", result.Email)
}
