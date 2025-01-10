package error

import "errors"

var (
	// error 400
	ErrNotFound                 = errors.New("not found")
	ErrBadRequest               = errors.New("bad request")
	ErrUsernameAlreadyUsed      = errors.New("username already used")
	ErrEmailAlreadyUsed         = errors.New("email already used")
	ErrIncorrectEmailOrPassword = errors.New("email or password incorrect")

	// error 500
	ErrInternalServerError = errors.New("internal server error")
	ErrConflict            = errors.New("conflict")
)