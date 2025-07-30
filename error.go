package apperr

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/protoadapt"
)

func NewError(cause error, reason, message string, code codes.Code, details ...protoadapt.MessageV1) *ErrorObject {
	return &ErrorObject{
		code:    code,
		reason:  reason,
		message: message,
		cause:   cause,
		details: details,
	}
}

type ErrorObject struct {
	code    codes.Code
	reason  string
	message string
	details []protoadapt.MessageV1
	cause   error
}

func (e *ErrorObject) Error() string {
	return e.message
}

func (e *ErrorObject) Unwrap() error {
	return e.cause
}

func (e *ErrorObject) Is(target error) bool {
	var be *ErrorObject
	if errors.As(target, &be) {
		return e.reason == be.reason
	}
	return false
}

func (e *ErrorObject) Reason() string {
	return e.reason
}

func (e *ErrorObject) Details() []protoadapt.MessageV1 {
	if e.details == nil {
		return nil
	}

	return e.details
}

func (e *ErrorObject) Code() codes.Code {
	return e.code
}
