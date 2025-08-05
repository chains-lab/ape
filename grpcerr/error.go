package grpcerr

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/protoadapt"
)

type ErrorMessage struct {
	// Reason is a unique identifier for the error type
	Reason string

	// Message  is a short, human-readable summary of the problem that SHOULD NOT change from occurrence to
	// occurrence of the problem, except for purposes of localization (analogue title in json api).
	Message string

	// Code is a gRPC status code that indicates the type of error
	Code codes.Code

	// Details is a list of additional details about the error, which can be used to provide more context
	Details []protoadapt.MessageV1
}

func (e *ErrorMessage) Error() string {
	return e.Message
}

func (e *ErrorMessage) Is(target error) bool {
	var be *ErrorMessage
	if errors.As(target, &be) {
		return e.Reason == be.Reason
	}
	return false
}
