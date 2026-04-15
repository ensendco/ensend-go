package ensend

import (
	"net/http"
	"os"

	"github.com/ensendco/ensend-go/config"
)

type Client struct {
	secret     string
	baseURL    string
	userAgent  string
	httpClient *http.Client

	transport http.RoundTripper

	SendApi *emailsService
}

func New(opts ...Option) *Client {
	secret := os.Getenv("ENSEND_PROJECT_SECRET_KEY")
	baseURL := os.Getenv("ENSEND_BASE_URL")
	if baseURL == "" {
		baseURL = config.BaseURL
	}

	c := &Client{
		secret:     secret,
		baseURL:    baseURL,
		userAgent:  config.UserAgent,
		transport:  http.DefaultTransport,
		httpClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(c)
	}

	// attach transport middleware stack
	c.httpClient.Transport = c.transport

	// register services
	c.SendApi = &emailsService{client: c}

	return c
}
