package persistence

import (
	"funding/internal/identity/domain"
)

type UserRow struct {
	ID         int           `db:"id"`
	UserID     domain.UserID `db:"user_id"`
	Name       string        `db:"name"`
	Email      string        `db:"email"`
	Password   string        `db:"password_hash"`
	Role       string        `db:"role"`
	Occupation string        `db:"occupation"`
	Avatar     string        `db:"avatar_file_name"`
	IsActive   bool          `db:"is_active"`
	IsVerified bool          `db:"is_verified"`
}
