package constant

import "errors"

var (
	ErrNotFound            = errors.New("record not found")
	ErrInternalServerError = errors.New("internal server error")
	ErrBadRequest          = errors.New("bad request")
	ErrConflict            = errors.New("conflict")
)
