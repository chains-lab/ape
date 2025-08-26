package problems

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/jsonapi"
)

func Conflict(details string) *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusConflict),
		Status: fmt.Sprintf("%d", http.StatusConflict),
		Detail: details,
		Meta: &map[string]any{
			"timestamp": time.Now().UTC(),
		},
	}
}
