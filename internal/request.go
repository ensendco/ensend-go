package internal

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Requester struct {
	Secret     string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
}

func (r *Requester) Do(
	ctx context.Context,
	method string,
	path string,
	body any,
	out any,
) (*http.Response, error) {

	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		r.BaseURL+path,
		&buf,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+r.Secret)
	req.Header.Set("User-Agent", r.UserAgent)
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if out != nil && resp.StatusCode < http.StatusBadRequest {
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
			return nil, err
		}
	}

	return resp, nil
}
