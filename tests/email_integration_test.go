package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	ensend "github.com/ensendco/ensend-go"
	"github.com/joho/godotenv"
)

func TestEmailsSendIntegration_DemoServer(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}

		if r.URL.Path != "/send" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message_id":"abc123","status":"sent"}`))
	}))
	defer server.Close()

	client := ensend.New()

	resp, err := client.Emails.Send(context.Background(), ensend.SendEmailRequestVars{
		Subject: "Demo integration test",
		Sender: ensend.Address{
			Name:    "Demo Sender",
			Address: "noreply@example.com",
		},
		Recipients: []ensend.Recipient[map[string]any]{
			{
				Name:    "Demo Recipient",
				Address: "recipient@example.com",
			},
		},
		Message: "This test validates SDK request/response wiring against a local demo server.",
	})

	if err != nil {
		t.Fatal(err)
	}

	if resp.MessageID != "abc123" {
		t.Fatalf("expected abc123 got %s", resp.MessageID)
	}
}

func TestEmailsSendIntegration_SMTPExpressCredentials(t *testing.T) {
	_ = godotenv.Load(".env")
	_ = godotenv.Load("../.env")

	apiKey := os.Getenv("ENSEND_PROJECT_SECRET_KEY")
	sender := os.Getenv("ENSEND_TEST_SENDER")
	recipient := os.Getenv("ENSEND_TEST_RECIPIENT")

	if apiKey == "" || sender == "" || recipient == "" {
		t.Skip("set ENSEND_API_KEY, ENSEND_TEST_SENDER, and ENSEND_TEST_RECIPIENT to run real SMTPExpress integration test")
	}

	baseURL := os.Getenv("ENSEND_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.smtpexpress.com"
	}

	client := ensend.New()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	resp, err := client.Emails.Send(ctx, ensend.SendEmailRequestVars{
		Subject: "SMTPExpress SDK integration test",
		Sender: ensend.Address{
			Name:    "SDK Integration",
			Address: sender,
		},
		Recipients: []ensend.Recipient[map[string]any]{
			{
				Name:    "SDK Recipient",
				Address: recipient,
			},
		},
		Message:      "This is an automated integration test from ensend_go_sdk.",
		ReplyAddress: sender,
	})
	if err != nil {
		t.Fatalf("real SMTPExpress send failed: %v", err)
	}

	if resp.MessageID == "" && (resp.Data == nil || resp.Data.Ref == "") {
		t.Fatalf("expected non-empty message identifier (message_id or data.ref), got response: %+v", resp)
	}
}
