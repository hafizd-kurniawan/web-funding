package postgres

import (
	"context"
	"funding/internal/domain/user"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

var _ user.Repository = (*UserRepository)(nil)

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(ctx context.Context, user *user.User) error {
	tx := r.db.MustBegin()
	defer tx.Commit()

	tx.MustExec(`
		INSERT INTO users (
		name, email, password, avatar_file_name, role, is_active, created_at, updated_at, deleted_at
		)
		VALUES (
		:name, :email, :password, :avatar_file_name, :role, :is_active, :created_at, :updated_at, :deleted_at)
	`, user)

	err := tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email user.Email) (*user.User, error) {
	return nil, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int) (*user.User, error) {
	return nil, nil
}
