package error

import "errors"

var (
	// error 400
	ErrNotFound                  = errors.New("not found")
	ErrBadRequest                = errors.New("bad request")
	ErrUsernameAlreadyUsed       = errors.New("username already used")
	ErrEmailAlreadyUsed          = errors.New("email already used")
	ErrIncorrectEmailOrPassword  = errors.New("email or password incorrect")
	ErrUnauthorized              = errors.New("unauthorized")
	ErrProfileAlreadyCreated     = errors.New("profile already created")
	ErrInvalidOTP                = errors.New("invalid otp")
	ErrInvalidPhone              = errors.New("invalid phone format")
	ErrInvalidTransactionType    = errors.New("invalid transaction type")
	ErrFailedToCreateTransaction = errors.New("failed to create transaction")
	ErrUnverifiedPhone           = errors.New("unverified phone")

	// error 500
	ErrInternalServerError = errors.New("internal server error")
	ErrConflict            = errors.New("conflict")
)
