package domain

import "github.com/golang-jwt/jwt/v5"

type TokenGenerator interface {
	GenerateToken(userID UserID) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
