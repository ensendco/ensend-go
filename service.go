package ensend

import (
	"context"
	"net/http"

	"github.com/ensendco/ensend_go_sdk/internal"
)

type emailsService struct {
	client *Client
}

// send is the shared, non-generic implementation.
func (s *emailsService) send(
	ctx context.Context,
	req any,
) (*SendEmailResponse, error) {

	r := internal.Requester{
		APIKey:     s.client.apiKey,
		BaseURL:    s.client.baseURL,
		UserAgent:  s.client.userAgent,
		HTTPClient: s.client.httpClient,
	}

	var out SendEmailResponse

	resp, err := r.Do(ctx, http.MethodPost, "/send", req, &out)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, parseAPIError(resp)
	}

	return &out, nil
}

// Send is a convenience method for requests with no custom variable types
// (uses map[string]any for recipient variables).
func (s *emailsService) Send(
	ctx context.Context,
	req SendEmailRequest[map[string]any],
) (*SendEmailResponse, error) {
	return s.send(ctx, req)
}

// Send is a top-level generic function for typed recipient variables.
// Usage:
//
//	resp, err := ensend.Send(ctx, client.Emails, ensend.SendEmailRequest[UserDetails]{...})
func Send[V any](
	ctx context.Context,
	svc *emailsService,
	req SendEmailRequest[V],
) (*SendEmailResponse, error) {
	return svc.send(ctx, req)
}
