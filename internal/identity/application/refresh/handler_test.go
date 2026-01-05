package refresh

import (
	"fmt"
	"funding/internal/identity/application/login"
	"funding/internal/identity/domain"
	jj "funding/internal/identity/infrastructure/jwt"
	"funding/internal/identity/infrastructure/persistence/postgres"
	conn "funding/shared/database/postgres"
	"testing"

	"funding/internal/identity/infrastructure/hasher"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type FakeUser struct {
	UserID       uuid.UUID
	RefreshToken uuid.UUID
}

var fakeUser FakeUser

func TestLoginUser(t *testing.T) {
	db := conn.SetupDB()
	repo := postgres.NewRepository(db)
	hasher := hasher.BcryptHasher{}
	accesToken := jj.NewJWTToken("secret")
	refreshToken := postgres.NewRefreshTokenRepository(db)
	handler := login.NewHandler(repo, hasher, accesToken, refreshToken)

	result, err := handler.Handle(login.Command{
		Email:    "ayo@gmail.commmmm",
		Password: "password123",
	})

	t.Log(result.UserID)
	fakeUser = FakeUser{
		UserID:       result.UserID,
		RefreshToken: result.RefreshToken,
	}

	assert.NoError(t, err)
	assert.Equal(t, "ayo@gmail.commmmm", result.Email)
}

func TestRefreshToken(t *testing.T) {
	db := conn.SetupDB()
	userRepo := postgres.NewRepository(db)
	refreshTokenRepo := postgres.NewRefreshTokenRepository(db)
	accesToken := jj.NewJWTToken("secret")

	// 1. Create and Save User
	userID := domain.UserID(uuid.New())
	email := fmt.Sprintf("test_refresh_%s@gmail.com", uuid.New().String())
	user, _, err := domain.RegisterUser(userID, "test_refresh", email)
	assert.NoError(t, err)
	err = userRepo.Save(user)
	assert.NoError(t, err)

	// 2. Create and Save Refresh Token
	refreshToken := domain.NewRefreshTokenID(userID)
	err = refreshTokenRepo.Save(refreshToken)
	assert.NoError(t, err)

	// 3. Test Handler
	handler := NewHandler(refreshTokenRepo, accesToken)
	result, err := handler.Handle(Command{
		UserID: refreshToken.ID,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, result.AccessToken)
	assert.NotEmpty(t, result.RefreshToken)
}
