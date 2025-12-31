package user

import "errors"

var (
	ErrInternalServerError  = errors.New("internal server error")
	ErrSQLError             = errors.New("database server failed to execute query")
	ErrTooManyRequests      = errors.New("too many requests")
	ErrUnauthorized         = errors.New("unauthorized")
	ErrInvalidToken         = errors.New("invalid token")
	ErrForbidden            = errors.New("forbidden")
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordInCorrect    = errors.New("password incorrect")
	ErrUsernameExist        = errors.New("username already exist")
	ErrNameNotValid         = errors.New("username already exist")
	ErrEmailExist           = errors.New("email already exist")
	ErrEmailNotValid        = errors.New("email not valid")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordInCorrect,
	ErrUsernameExist,
	ErrPasswordDoesNotMatch,
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}

func IsUserError(err error) bool {
	for _, e := range UserErrors {
		if errors.Is(err, e) {
			return true
		}
	}
	return false
}
