package util

import "fmt"

type ErrorType int64

const (
	ErrNotFound            ErrorType = 404
	ErrAborted             ErrorType = 505
	ErrorUnauthorized      ErrorType = 401
	ErrBadRequest          ErrorType = 400
	ErrParseJson           ErrorType = 400
	ErrInternalServerError ErrorType = 500
)

type MainError struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
}

func (e *MainError) Error() string {
	return fmt.Sprintf("%d-%s", e.Type, e.Message)
}

var (
	UnauthorizedError = &MainError{
		Type:    ErrorUnauthorized,
		Message: "unauthorized",
	}
	InternalServerError = &MainError{
		Type:    ErrInternalServerError,
		Message: "internal server error",
	}
	BadRequestError = &MainError{
		Type:    ErrBadRequest,
		Message: "bad request error",
	}
)
