package ensend

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var payload struct {
		Error   string `json:"error"`
		Message string `json:"message"`
		Detail  string `json:"detail"`
	}
	_ = json.Unmarshal(body, &payload)

	msg := strings.TrimSpace(payload.Error)
	if msg == "" {
		msg = strings.TrimSpace(payload.Message)
	}
	if msg == "" {
		msg = strings.TrimSpace(payload.Detail)
	}
	if msg == "" {
		msg = strings.TrimSpace(string(body))
	}

	return &APIError{
		StatusCode: resp.StatusCode,
		Message:    msg,
		RequestID:  resp.Header.Get("X-Request-ID"),
	}
}
