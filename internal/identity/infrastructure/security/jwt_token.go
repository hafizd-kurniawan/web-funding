package security

import (
	"errors"
	"funding/internal/identity/domain"

	"github.com/golang-jwt/jwt/v5"
)

type JWTToken struct {
	secret string
}

var _ domain.TokenGenerator = JWTToken{}

func NewJWTToken(secret string) JWTToken {
	return JWTToken{
		secret: secret,
	}
}

func (t JWTToken) GenerateToken(userID domain.UserID) (string, error) {
	claim := jwt.MapClaims{
		"user_id": userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(t.secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (t JWTToken) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(t.secret), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}
