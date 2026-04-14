package main

import (
	"context"
	"fmt"

	ensend "github.com/ensendco/ensend-go"
)

func main() {
	client := ensend.New("api-key",
		ensend.WithMiddleware(ensend.LoggingMiddleware),
	)

	type UserDetails struct {
		FirstName string
	}

	resp, err := ensend.Send(context.Background(), client.Emails, ensend.SendEmailRequest[UserDetails]{
		Subject: "{{firstName}}, Welcome to Ensend",
		Sender: ensend.Address{
			Name:    "Sender Name",
			Address: "noreply@ensend.me",
		},
		Recipients: []ensend.Recipient[UserDetails]{
			{
				Name:    "Ensend Labs",
				Address: "tenotea@ensend.me",
				Variables: UserDetails{
					FirstName: "Ensend",
				},
			},
		},
		Message:      "{{firstName}}, this is a test message.",
		ReplyAddress: "noreply@ensend.me",
		Attachments: []ensend.Attachment{
			{
				Name: "File.pdf",
				URL:  "https://link.to/file.pdf",
			},
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Sent:", resp.MessageID)
}
