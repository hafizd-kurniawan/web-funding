package postgres

import (
	"fmt"
	"funding/internal/identity/domain"

	"funding/internal/identity/infrastructure/persistence"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

var _ domain.UserRepository = Repository{}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) Save(user *domain.User) error {
	query := `
		INSERT INTO users (
			user_id, name, occupation, email, password_hash, avatar_file_name, role, is_active, is_verified
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.Exec(
		query,
		user.GetUUID(),
		user.GetName(),
		user.GetOccupation(),
		user.GetEmail(),
		user.GetPassword().Value(),
		user.GetAvatar(),
		user.GetRole(),
		user.IsUserActive(),
		user.IsUserVerified(),
	)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (r Repository) FindByEmail(email string) (*domain.User, error) {
	query := `
		SELECT id, user_id, name, occupation, email, password_hash, avatar_file_name, role, is_active, is_verified
		FROM users
		WHERE email = $1
	`
	var user persistence.UserRow
	err := r.db.Get(&user, query, email)
	if err != nil {
		return nil, err
	}

	return domain.RehydrateUser(user.ID, user.UserID, user.Name, user.Email, user.Password, user.Role, user.Occupation, user.Avatar, user.IsActive, user.IsVerified), nil

}

func (r Repository) FindByID(id int) (*domain.User, error) {
	query := `
		SELECT id, name, occupation, email, password_hash, avatar_file_name, role, is_active, is_verified
		FROM users
		WHERE id = $1
	`
	var user domain.User
	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
