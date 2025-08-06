package ape

import (
	"errors"

	"google.golang.org/grpc/status"
)

type Error struct {
	// id unique error identifier
	// in uppercase format like "ADMIN_CAN_NOT_DELETE_SELF"
	id string

	// internal error which caused this error
	Cause error

	//Response error
	Response *status.Status
}

func (e *Error) Error() string {
	return e.id
}

func (e *Error) Is(target error) bool {
	var be *Error
	if errors.As(target, &be) {
		return e.id == be.id
	}
	return false
}

func Declare(id string) *Error {
	return &Error{
		id: id,
	}
}

func (e *Error) Raise(cause error, response *status.Status) *Error {
	return &Error{
		id:       e.id,
		Cause:    cause,
		Response: response,
	}
}
