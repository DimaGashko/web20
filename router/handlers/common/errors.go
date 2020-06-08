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
	appErr, isAppErr := err.(*AppError)
	if isAppErr {
		return appErr
	}
	return &AppError{err, code}
}

func New(code int) *AppError {
	return MakeError("", code)
}

func New403() *AppError {
	return New(http.StatusForbidden)
}

func New404() *AppError {
	return New(http.StatusNotFound)
}

func New500() *AppError {
	return New(http.StatusInternalServerError)
}

func MakeError(msg string, code int) *AppError {
	return WrapError(errors.New(msg), code)
}

func GetErrorCode(err error) int {
	appErr, ok := err.(*AppError)
	if !ok {
		return http.StatusInternalServerError
	}
	return appErr.Code
}
