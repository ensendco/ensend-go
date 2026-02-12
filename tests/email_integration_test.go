package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	ensend "github.com/xpanvictor/ensend_go_sdk"
)

func TestEmailsSendIntegration(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/emails/send" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message_id":"abc123","status":"sent"}`))
	}))
	defer server.Close()

	client := ensend.New("test-key",
		ensend.WithBaseURL(server.URL),
	)

	resp, err := client.Emails.Send(context.Background(), ensend.SendEmailRequest[map[string]any]{
		Subject: "Test",
	})

	if err != nil {
		t.Fatal(err)
	}

	if resp.MessageID != "abc123" {
		t.Fatalf("expected abc123 got %s", resp.MessageID)
	}
}