package response

import (
	"errors"
	"net/http"
)

// error general
var (
	ErrNotFound = errors.New("not found")
)

var (
	ErrEmailRequired = errors.New("email is required")
	ErrEmailInvalid  = errors.New("email is invalid")

	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password must have at least 6 characters")
	ErrPasswordNotMatch      = errors.New("password is not matched")

	ErrAuthIsNotExists  = errors.New("auth is not exists")
	ErrEmailAlreadyUsed = errors.New("email is already used")
)

type Error struct {
	Message  string
	Code     string
	HttpCode int
}

func NewError(msg string, code string, httpCode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral    = NewError("general error", "99999", http.StatusInternalServerError)
	ErrorBadRequest = NewError("bad request", "40000", http.StatusBadRequest)
)

var (
	ErrorEmailRequired         = NewError(ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorEmailInvalid          = NewError(ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
	ErrorPasswordRequired      = NewError(ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)
	ErrorPasswordNotMatch      = NewError(ErrPasswordNotMatch.Error(), "40101", http.StatusNotFound)
	ErrorAuthIsNotExists       = NewError(ErrAuthIsNotExists.Error(), "40401", http.StatusConflict)
	ErrorEmailAlreadyUsed      = NewError(ErrEmailAlreadyUsed.Error(), "40901", http.StatusUnauthorized)
	ErrorNotFound              = NewError(ErrNotFound.Error(), "40400", http.StatusNotFound)
)

var (
	ErrorMapping = map[string]Error{
		ErrNotFound.Error():              ErrorNotFound,
		ErrEmailRequired.Error():         ErrorEmailRequired,
		ErrEmailInvalid.Error():          ErrorEmailInvalid,
		ErrPasswordRequired.Error():      ErrorPasswordRequired,
		ErrPasswordInvalidLength.Error(): ErrorPasswordInvalidLength,
		ErrPasswordNotMatch.Error():      ErrorPasswordNotMatch,
		ErrAuthIsNotExists.Error():       ErrorAuthIsNotExists,
		ErrEmailAlreadyUsed.Error():      ErrorEmailAlreadyUsed,
	}
)
