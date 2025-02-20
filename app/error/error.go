package error

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Error struct {
	Code     int
	HttpCode int
	Message  string
}

func NewError(message string) *Error {
	err := &Error{}
	switch message {
	case ErrBadRequest.Error():
		err.Code = 4001
		err.HttpCode = http.StatusBadRequest
		err.Message = message
	case ErrNotFound.Error():
		err.Code = 4002
		err.HttpCode = http.StatusNotFound
		err.Message = message
	case ErrUsernameAlreadyUsed.Error():
		err.Code = 4003
		err.HttpCode = http.StatusBadRequest
		err.Message = message
	case ErrEmailAlreadyUsed.Error():
		err.Code = 4004
		err.HttpCode = http.StatusBadRequest
		err.Message = message
	case ErrIncorrectEmailOrPassword.Error():
		err.Code = 4005
		err.HttpCode = http.StatusBadRequest
		err.Message = message
	case ErrUnauthorized.Error():
		err.Code = 4006
		err.HttpCode = http.StatusUnauthorized
		err.Message = message
	case ErrProfileAlreadyCreated.Error():
		err.Code = 4007
		err.HttpCode = http.StatusBadRequest
		err.Message = message
	case ErrInvalidOTP.Error():
		err.Code = 4008
		err.HttpCode = http.StatusBadRequest
		err.Message = message
	case ErrInvalidPhone.Error():
		err.Code = 4009
		err.HttpCode = http.StatusBadRequest
		err.Message = message
	case ErrInvalidTransactionType.Error():
		err.Code = 4010
		err.HttpCode = http.StatusBadRequest
		err.Message = message
	case ErrFailedToCreateTransaction.Error():
		err.Code = 4011
		err.HttpCode = http.StatusBadRequest
		err.Message = message
	case ErrUnverifiedPhone.Error():
		err.Code = 4012
		err.HttpCode = http.StatusBadRequest
		err.Message = message
	case ErrThereIsPendingTransaction.Error():
		err.Code = 4013
		err.HttpCode = http.StatusBadRequest
		err.Message = message
	default:
		err.Code = 5001
		err.HttpCode = http.StatusInternalServerError
		err.Message = ErrInternalServerError.Error()
	}

	return err
}

func ErrorResponse(err error, c *gin.Context) {
	newError := NewError(err.Error())

	// if validationErrs, ok := err.(validator.ValidationErrors); ok {

	// }

	switch err.(type) {
	case validator.ValidationErrors:
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    4000,
			"message": err.Error(),
		})
	case *json.UnmarshalTypeError,
		*json.UnsupportedTypeError,
		*json.UnsupportedValueError:
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    4000,
			"message": err.Error(),
		})
	default:
		c.JSON(newError.HttpCode, gin.H{
			"code":    newError.Code,
			"message": newError.Message,
		})
	}

}
