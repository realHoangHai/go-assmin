package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
}

func NewErrorResponse(status int, root error, msg string, args ...interface{}) *AppError {
	return &AppError{
		StatusCode: status,
		RootErr:    root,
		Message:    fmt.Sprintf(msg, args...),
	}
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func NewCustomError(status int, root error, msg string, args ...interface{}) *AppError {
	if root != nil {
		return NewErrorResponse(status, root, msg)
	}
	return NewErrorResponse(status, errors.New(msg), msg)
}

func ErrInvalidRequest(root error, msg string, args ...interface{}) *AppError {
	return NewCustomError(http.StatusBadRequest, root, msg, args...)
}

func ErrInternal(root error, msg string, args ...interface{}) *AppError {
	return NewErrorResponse(http.StatusInternalServerError, root, msg, args...)
}

var (
	ErrInvalidToken     = NewCustomError(http.StatusUnauthorized, nil, "invalid signature")
	ErrUnauthoried      = NewCustomError(http.StatusUnauthorized, nil, "unauthorized")
	ErrNotFound         = NewCustomError(http.StatusNotFound, nil, "not found")
	ErrTooManyRequests  = NewCustomError(http.StatusTooManyRequests, nil, "too many requests")
	ErrMethodNotAllowed = NewCustomError(http.StatusMethodNotAllowed, nil, "method not allowed")
	ErrExpiredToken     = ErrInvalidRequest(nil, "token has expired")
)
