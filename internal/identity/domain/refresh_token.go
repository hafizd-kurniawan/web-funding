package domain

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID        uuid.UUID `db:"id"`
	UserID    UserID    `db:"user_id"`
	ExpiresAt time.Time `db:"expires_at"`
	Revoked   bool      `db:"revoked"`
	RevokedAt time.Time `db:"revoked_at"`
	CreatedAt time.Time `db:"created_at"`
}

type RefreshTokenRepo interface {
	Save(RefreshToken) error
	FindByID(uuid.UUID) (RefreshToken, error)
	Revoke(uuid.UUID) error
}

func NewRefreshTokenID(userID UserID) RefreshToken {
	return RefreshToken{
		ID:        uuid.New(),
		UserID:    userID,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
		Revoked:   false,
	}
}

func (r RefreshToken) IsExpired() bool {
	return time.Now().After(r.ExpiresAt)
}
