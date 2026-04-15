package ensend

type Address struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"email"`
}

type Recipient[V any] struct {
	Name      string `json:"name,omitempty"`
	Address   string `json:"email"`
	Variables V      `json:"variables,omitempty"`
}

type TemplateRef[V any] struct {
	ID        string `json:"id"`
	Variables V      `json:"variables,omitempty"`
}

type Attachment struct {
	Name    string `json:"name"`
	URL     string `json:"url,omitempty"`
	Content string `json:"content,omitempty"`
}

type SendEmailRequestVars = SendEmailRequest[map[string]any, map[string]any]
type SendEmailRequest[V any, K any] struct {
	Subject      string          `json:"subject"`
	Sender       Address         `json:"sender"`
	Recipients   []Recipient[V]  `json:"recipients"`
	Message      string          `json:"message,omitempty"`
	Template     *TemplateRef[K] `json:"template,omitempty"`
	ReplyAddress string          `json:"replyAddress,omitempty"`
	Attachments  []Attachment    `json:"attachments,omitempty"`
}

type SendEmailResponse struct {
	MessageID  string                 `json:"message_id,omitempty"`
	Status     string                 `json:"status,omitempty"`
	Message    string                 `json:"message,omitempty"`
	StatusCode int                    `json:"statusCode,omitempty"`
	Data       *SendEmailResponseData `json:"data,omitempty"`
}

type SendEmailResponseData struct {
	Ref string `json:"ref,omitempty"`
}
