package requests_models

type RequestType string

const (
	TypeMessage RequestType = "Message" // Send Message to Chat
	TypeConfig  RequestType = "Config"  // Change Game Config (only Host)
)

type RequestEvent struct {
	Type    RequestType    `json:"type" validate:"required,min=3,max=20"`
	Message string         `json:"message" validate:"required,min=3,max=50"`
	Changes map[string]any `json:"changes"`
}
