package ensend

import "net/http"

type Middleware func(http.RoundTripper) http.RoundTripper

func WithMiddleware(m Middleware) Option {
	return func(c *Client) {
		c.transport = m(c.transport)
	}
}