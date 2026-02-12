package ensend

import (
	"net/http"
)

func LoggingMiddleware(next http.RoundTripper) http.RoundTripper {
	return roundTripperFunc(func(req *http.Request) (*http.Response, error) {

		println("Sent:", req.Method, req.URL.String())

		resp, err := next.RoundTrip(req)

		if err != nil {
			println("request failed:", err.Error())
			return nil, err
		}

		println("Recv status:", resp.StatusCode)

		return resp, nil
	})
}