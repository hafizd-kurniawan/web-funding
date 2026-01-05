package login

import "github.com/google/uuid"

type Result struct {
	Name         string
	UserID       uuid.UUID
	Email        string
	Token        string
	RefreshToken uuid.UUID
}
