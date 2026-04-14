# Ensend Go SDK

Go client for sending emails through SMTPExpress/Ensend.

## Install

```bash
go get github.com/xpanvictor/ensend_go_sdk
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"

	ensend "github.com/xpanvictor/ensend_go_sdk"
)

func main() {
	client := ensend.New("YOUR_API_KEY",
		ensend.WithBaseURL("https://api.smtpexpress.com"),
	)

	resp, err := client.Emails.Send(context.Background(), ensend.SendEmailRequest[map[string]any]{
		Subject: "My first email using the SDK",
		Message: "Hello from Go",
		Sender: ensend.Address{
			Name:    "SMTP Express User",
			Address: "documentation@ensend.me", // serialized as sender.email
		},
		Recipients: []ensend.Recipient[map[string]any]{
			{
				Address: "recipient@example.com", // serialized as recipients[].email
			},
		},
	})
	if err != nil {
		panic(err)
	}

	ref := resp.MessageID
	if ref == "" && resp.Data != nil {
		ref = resp.Data.Ref
	}

	fmt.Println("Sent reference:", ref)
}
```

## Request Shape

- `sender.Address` is sent as JSON field `sender.email`
- `recipient.Address` is sent as JSON field `recipients[].email`
- Endpoint used by this SDK: `POST /send`

## Response Shape

Success responses may return either:

- `message_id` (mapped to `SendEmailResponse.MessageID`)
- `data.ref` (mapped to `SendEmailResponse.Data.Ref`)

Check both fields if you need a send reference.

## Options

- `WithBaseURL(url string)`
- `WithTimeout(d time.Duration)`
- `WithHTTPClient(client *http.Client)`
- `WithUserAgent(agent string)`
- `WithMiddleware(m Middleware)`

Example middleware:

```go
client := ensend.New("YOUR_API_KEY",
	ensend.WithMiddleware(ensend.LoggingMiddleware),
)
```

## Integration Tests

The credentials-based integration test reads these env vars:

- `ENSEND_API_KEY`
- `ENSEND_TEST_SENDER`
- `ENSEND_TEST_RECIPIENT`
- `ENSEND_BASE_URL` (optional, defaults in test to `https://api.smtpexpress.com`)

Run tests:

```bash
go test ./tests -v
```