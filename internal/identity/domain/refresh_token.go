package domain

import (
	"time"

	"github.com/google/uuid"
)

type RefreshTokenRepo interface {
	Save(RefreshToken) error
	FindByID(uuid.UUID) (RefreshToken, error)
	Revoke(uuid.UUID) error
}

type RefreshToken struct {
	ID        uuid.UUID
	UserID    UserID
	ExpiresAt time.Time
	Revoked   bool
}

func (r RefreshToken) IsExpired() bool {
	return r.ExpiresAt.After(time.Now())
}
