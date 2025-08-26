package problems

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/jsonapi"
)

func TooManyRequests() *jsonapi.ErrorObject {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusTooManyRequests),
		Status: fmt.Sprintf("%d", http.StatusTooManyRequests),
		Meta: &map[string]interface{}{
			"timestamp": time.Now().UTC(),
		},
	}
}
