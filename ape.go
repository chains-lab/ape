package ape

import (
	"errors"
)

type Error struct {
	// ID unique error identifier
	// in uppercase format like "ADMIN_CAN_NOT_DELETE_SELF"
	ID string

	// internal error which caused this error
	Cause error
}

func (e *Error) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Cause != nil {
		return e.Cause.Error()
	}
	return e.ID
}

func (e *Error) Is(target error) bool {
	var be *Error
	if errors.As(target, &be) {
		return e.ID == be.ID
	}
	return false
}

func (e *Error) Raise(cause error) error {
	return &Error{
		ID:    e.ID,
		Cause: cause,
	}
}

func DeclareError(ID string) *Error {
	return &Error{
		ID: ID,
	}
}
