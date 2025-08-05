package grpcerr

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/protoadapt"
)

type ErrorMessage struct {
	// Code is a gRPC status code that indicates the type of error
	Code codes.Code

	// Reason is an application-specific error code, expressed as a string value.
	Reason string

	// The Title is a short, human-readable summary of the problem that SHOULD NOT change from occurrence to
	// occurrence of the problem, except for purposes of localization (analogue title in json api).
	Title string

	// RequestID is a unique identifier for the request, which can be used for tracing and debugging
	RequestID string

	// Message is a list of additional details about the error, which can be used to provide more context
	Message []protoadapt.MessageV2
}

func (e *ErrorMessage) Error() string {
	return e.Title
}

func (e *ErrorMessage) Is(target error) bool {
	var be *ErrorMessage
	if errors.As(target, &be) {
		return e.Reason == be.Reason
	}
	return false
}
