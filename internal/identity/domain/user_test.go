package domain

import (
	"errors"
	"funding/internal/identity/application"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type fakeHasher struct{}


// type fakeUserRepository struct{}
//
// func (fakeUserRepository) Save(user *User) error {
// 	return nil
// }
//
// func (fakeUserRepository) FindByEmail(email string) (*User, error) {
// 	return nil, nil
// }
//
// func (fakeUserRepository) FindByID(id int) (*User, error) {
// 	return nil, nil
// }

type fakeTokenGenerator struct{}

func (fakeTokenGenerator) GenerateToken(userID UserID) (string, error) {
	return "", nil
}

func (fakeTokenGenerator) ValidateToken(token string) (*jwt.Token, error) {
	return nil, nil
}

func (fakeHasher) Hash(plain string) ([]byte, error) {
	if plain == "" {
		return []byte{}, errors.New("password is required")
	}
	hashed := "hashed-" + plain
	return []byte(hashed), nil
}

func (fakeHasher) Compare(plain, hashed string) bool {
	return hashed == "hashed-"+plain
}

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		testCase string
		name     string
		email    string
		password string
		wantErr  bool
	}{
		{"emptyName", "", "ayo@gmail.com", "password", true},
		{"invalidLengthName", "a", "ayo@gmail.com", "password", true},
		{"validName", "ayo", "ayo@gmail.com", "password", false},
		{"invalidEmail", "ayo", "ayo", "password", true},
		{"validEmail", "ayo", "ayo@gmail.com", "password", false},
		{"emptyEmail", "ayo", "", "password", true},
	}

	for _, tt := range tests {
		tt := tt // IMPORTANT
		t.Run(tt.testCase, func(t *testing.T) {
			t.Parallel()

			userID := UserID(uuid.New())

			_, _, err := RegisterUser(userID, tt.name, tt.email)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestChangePassword(t *testing.T) {
	tests := []struct {
		testCase string
		password string
		wantErr  bool
	}{
		{"emptyPassword", "", true},
		{"invalidLengthPassword", "a", true},
		{"validPassword", "password", false},
	}

	hasher := fakeHasher{}

	for _, tt := range tests {
		tt := tt // IMPORTANT
		t.Run(tt.testCase, func(t *testing.T) {
			t.Parallel()

			userID := UserID(uuid.New())
			user, _, err := RegisterUser(userID, "ayo", "ayo@gmail.com")
			if err != nil {
				t.Fatal(err)
			}

			password, err := NewPassword(tt.password, hasher)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			err = user.ChangePassword(password)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestChangeAvatar(t *testing.T) {
	tests := []struct {
		testCase string
		avatar   string
		wantErr  bool
	}{
		{"emptyAvatar", "", true},
		{"invalidLengthAvatar", "a", true},
		{"validAvatar", "avatar", false},
	}

	for _, tt := range tests {
		tt := tt // IMPORTANT
		t.Run(tt.testCase, func(t *testing.T) {
			t.Parallel()

			userID := UserID(uuid.New())
			user, _, err := RegisterUser(userID, "ayo", "ayo@gmail.com")
			if err != nil {
				t.Fatal(err)
			}

			err = user.ChangeAvatar(tt.avatar)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	tests := []struct {
		testCase string
		email    string
		password string
		wantErr  bool
	}{
		{"emptyEmail", "", "password", true},
		{"invalidEmail", "ayo", "password", true},
		{"validEmail", "ayo@gmail.com", "password", false},
		{"emptyPassword", "", "password", true},
		{"invalidPassword", "ayo@gmail.com", "", true},
		{"validPassword", "ayo@gmail.com", "password", false},
	}

	hasher := fakeHasher{}
	var _ domain.UserRepository = fakeUserRepository{}
	registerUser := application.RegisterUser{
		Repo:          fakeUserRepository{},
	}

	for _, tt := range tests {
		tt := tt // IMPORTANT
		t.Run(tt.testCase, func(t *testing.T) {
			t.Parallel()

			userID := UserID(uuid.New())
			user, _, err := RegisterUser(userID, "ayo", "ayo@gmail.com")
			if err != nil {
				t.Fatal(err)
			}

			password, err := NewPassword(tt.password, hasher)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			err = user.ChangePassword(password)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			token, err := token.GenerateToken(userID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			result, err := LoginUser(userID, tt.email, password)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, result.Email, tt.email)
			}
		})
	}
}
