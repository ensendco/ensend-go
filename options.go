package ensend

import (
	"net/http"
	"time"
)

type Option func(*Client)

func WithProjectSecret(secret string) Option {
	return func(c *Client) {
		c.secret = secret
	}
}

func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = d
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.httpClient = client
	}
}

func WithUserAgent(agent string) Option {
	return func(c *Client) {
		c.userAgent = agent
	}
}
