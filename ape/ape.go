package ape

import "errors"

type AppError struct {
	//code of error, should be unique for each error type
	// in upercase, e.g. "USER_IS_NOT_ADMIN"
	code string

	//internal reason for the error, should not be used in user messages
	cause error
}

func (e *AppError) Error() string {
	return e.code
}

func (e *AppError) Unwrap() error {
	if e.cause != nil {
		return e.cause
	}
	return nil
}

func (e *AppError) Is(target error) bool {
	var be *AppError
	if errors.As(target, &be) {
		return e.code == be.code
	}
	return false
}

func Create(code string) *AppError {
	return &AppError{
		code: code,
	}
}

func Raise(code string, cause error) *AppError {
	return &AppError{
		code:  code,
		cause: cause,
	}
}
