package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	Err  error
	Code int
}

func (e AppError) Error() string {
	return e.Err.Error()
}

func WrapError(err error, code int) *AppError {
	return &AppError{err, code}
}

func MakeError(msg string, code int) *AppError {
	if msg == "" {
		msg = http.StatusText(code)
	}

	return WrapError(errors.New(msg), code)
}
