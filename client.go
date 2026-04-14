package ensend

import (
	"net/http"

	"github.com/ensendco/ensend-go"
)

type Client struct {
	apiKey     string
	baseURL    string
	userAgent  string
	httpClient *http.Client

	transport http.RoundTripper

	Emails *emailsService
}

func New(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey:    apiKey,
		baseURL:   config.BaseURL,
		userAgent: config.UserAgent,

		transport:  http.DefaultTransport,
		httpClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(c)
	}

	// attach transport middleware stack
	c.httpClient.Transport = c.transport

	// register services
	c.Emails = &emailsService{client: c}

	return c
}
