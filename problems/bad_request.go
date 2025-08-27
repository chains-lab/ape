package problems

import (
	"fmt"
	"io"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/jsonapi"
	"github.com/pkg/errors"
)

// BadRequester is an error that indicates bad request.
type BadRequester interface {
	BadRequest() map[string]error
}

func BadRequest(err error) []*jsonapi.ErrorObject {
	cause := errors.Cause(err)
	if cause == io.EOF {
		return []*jsonapi.ErrorObject{
			{
				Title:  http.StatusText(http.StatusBadRequest),
				Status: fmt.Sprintf("%d", http.StatusBadRequest),
				Code:   "BAD_REQUEST",
				Detail: "Request body were expected",
				Meta: &map[string]any{
					"timestamp": time.Now().UTC(),
				},
			},
		}
	}

	switch cause := cause.(type) {
	case validation.Errors:
		return toJsonapiErrors(cause)
	case BadRequester:
		return toJsonapiErrors(cause.BadRequest())
	default:
		return []*jsonapi.ErrorObject{
			{
				Title:  http.StatusText(http.StatusBadRequest),
				Status: fmt.Sprintf("%d", http.StatusBadRequest),
				Code:   "BAD_REQUEST",
				Detail: "Your request was invalid in some way",
				Meta: &map[string]any{
					"timestamp": time.Now().UTC(),
				},
			},
		}
	}
}

func toJsonapiErrors(m map[string]error) []*jsonapi.ErrorObject {
	errs := make([]*jsonapi.ErrorObject, 0, len(m))
	for key, value := range m {
		errs = append(errs, &jsonapi.ErrorObject{
			Title:  http.StatusText(http.StatusBadRequest),
			Status: fmt.Sprintf("%d", http.StatusBadRequest),
			Code:   "BAD_REQUEST",
			Meta: &map[string]interface{}{
				"field":     key,
				"error":     value.Error(),
				"timestamp": time.Now().UTC(),
			},
		})
	}
	return errs
}

func InvalidParameter(parameter string, reason error) *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusBadRequest),
		Status: fmt.Sprintf("%d", http.StatusBadRequest),
		Code:   "BAD_REQUEST",
		Detail: fmt.Sprintf("Invalid parameter: %s", parameter),
		Meta: &map[string]interface{}{
			"parameter": parameter,
			"reason":    reason.Error(),
			"timestamp": time.Now().UTC(),
		},
	}
}

func InvalidPointer(pointer string, reason error) *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusBadRequest),
		Status: fmt.Sprintf("%d", http.StatusBadRequest),
		Code:   "BAD_REQUEST",
		Detail: fmt.Sprintf("Invalid pointer: %s", pointer),
		Meta: &map[string]interface{}{
			"pointer":   pointer,
			"reason":    reason.Error(),
			"timestamp": time.Now().UTC(),
		},
	}
}
