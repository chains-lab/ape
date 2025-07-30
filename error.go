package apperr

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/protoadapt"
)

type ErrorObject struct {
	Cause   error
	Reason  string
	Message string
	Code    codes.Code
	Details []protoadapt.MessageV1
}

func (e *ErrorObject) Error() string {
	return e.Message
}

func (e *ErrorObject) Unwrap() error {
	return e.Cause
}

func (e *ErrorObject) Is(target error) bool {
	var be *ErrorObject
	if errors.As(target, &be) {
		return e.Reason == be.Reason
	}
	return false
}
