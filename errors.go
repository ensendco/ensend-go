package ensend

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int
	Message    string
	RequestID  string
}

func (e *APIError) Error() string {
	return fmt.Sprintf(
		"smtp api error: status=%d message=%q request_id=%s",
		e.StatusCode,
		e.Message,
		e.RequestID,
	)
}

func parseAPIError(resp *http.Response) error {
	var payload struct {
		Error string `json:"error"`
	}

	json.NewDecoder(resp.Body).Decode(&payload)

	return &APIError{
		StatusCode: resp.StatusCode,
		Message:    payload.Error,
		RequestID:  resp.Header.Get("X-Request-ID"),
	}
}