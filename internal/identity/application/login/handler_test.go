package login

import (
	"testing"

	"funding/internal/identity/infrastructure/hasher"
	jj "funding/internal/identity/infrastructure/jwt"
	"funding/internal/identity/infrastructure/persistence/postgres"
	conn "funding/shared/database/postgres"

	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	db := conn.SetupDB()
	repo := postgres.NewRepository(db)
	hasher := hasher.BcryptHasher{}
	accesToken := jj.NewJWTToken("secret")
	refreshToken := postgres.NewRefreshTokenRepository(db)
	handler := NewHandler(repo, hasher, accesToken, refreshToken)

	result, err := handler.Handle(Command{
		Email:    "ayo@gmail.commmmm",
		Password: "password123",
	})

	assert.NoError(t, err)
	assert.Equal(t, "ayo@gmail.commmmm", result.Email)
}
