package utils

import "errors"

var (
	ErrNotFound            = errors.New("not found")
	ErrInternalServerError = errors.New("internal server error")
)

func IsNotFound(err error) bool {
	return err == ErrNotFound
}
