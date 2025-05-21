package requests_models

type RequestType string
type ResponseType string

const (
	StatusMessage RequestType = "Message"
	StatusConfig  RequestType = "Config"
)

const (
	TypeError   ResponseType = "Error"
	TypeMessage ResponseType = "Message"
	TypeConfig  ResponseType = "Config"
)

type RequestEvent struct {
	Type    RequestType    `json:"type"`
	Message string         `json:"message"`
	Changes map[string]any `json:"changes"`
}

type ResponseEvent struct {
	Success bool         `json:"success"`
	Type    ResponseType `json:"type"`
	Message string       `json:"message"`
	Content string       `json:"content"`
}
