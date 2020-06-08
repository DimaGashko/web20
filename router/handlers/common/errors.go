package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	Err  error `json:"error"`
	Code int   `json:"code"`
}

func (e AppError) Error() string {
	return e.Err.Error()
}

func WrapError(err error, code int) *AppError {
	appErr, isAppErr := err.(*AppError)
	if isAppErr {
		return appErr
	}
	return &AppError{err, code}
}

func New(msg string, code int) *AppError {
	return MakeError(msg, code)
}

func New403(msg string) *AppError {
	return New(msg, http.StatusForbidden)
}

func New404(msg string) *AppError {
	return New(msg, http.StatusNotFound)
}

func New500(msg string) *AppError {
	return New(msg, http.StatusInternalServerError)
}

func MakeError(msg string, code int) *AppError {
	if msg == "" {
		msg = http.StatusText(code)
	}
	return WrapError(errors.New(msg), code)
}

func GetErrorCode(err error) int {
	appErr, ok := err.(*AppError)
	if !ok {
		return http.StatusInternalServerError
	}
	return appErr.Code
}
