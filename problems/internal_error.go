package problems

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/jsonapi"
)

func InternalError(details, requestID string) *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusInternalServerError),
		Status: fmt.Sprintf("%d", http.StatusInternalServerError),
		Detail: details,
		Meta: &map[string]any{
			"timestamp":  time.Now().UTC(),
			"request_id": requestID,
		},
	}
}
