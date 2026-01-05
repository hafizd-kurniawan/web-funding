package postgres

import (
	"funding/internal/identity/domain"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RefreshTokenRepository struct {
	db *sqlx.DB
}

var _ domain.RefreshTokenRepo = RefreshTokenRepository{}

func NewRefreshTokenRepository(db *sqlx.DB) RefreshTokenRepository {
	return RefreshTokenRepository{
		db: db,
	}
}

func (r RefreshTokenRepository) Save(refreshToken domain.RefreshToken) error {
	query := `INSERT INTO refresh_tokens (
		id, user_id, expires_at, revoked, revoked_at, created_at
	) VALUES ($1, $2, $3, $4, NULL, NOW())`

	_, err := r.db.Exec(
		query,
		refreshToken.ID,
		refreshToken.UserID.String(),
		refreshToken.ExpiresAt,
		refreshToken.Revoked,
	)
	return err
}

func (r RefreshTokenRepository) FindByID(id uuid.UUID) (domain.RefreshToken, error) {
	query := `SELECT id, user_id, expires_at, revoked, created_at FROM refresh_tokens WHERE id = $1`
	var refreshToken domain.RefreshToken
	err := r.db.Get(&refreshToken, query, id.String())
	if err != nil {
		return refreshToken, err
	}
	return refreshToken, nil
}

func (r RefreshTokenRepository) Revoke(id uuid.UUID) error {
	query := `UPDATE refresh_tokens SET revoked = true, revoked_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
