package requests_models

type RequestAction string

const (
	TypeMessage RequestAction = "Message" // Send Message to Chat
	TypeConfig  RequestAction = "Config"  // Change Game Config (only Host)
	TypePlay    RequestAction = "Play"    // Play Game (only Host)
	TypeKick    RequestAction = "Kick"    // Kick Player (only Host)
)

type RequestEvent struct {
	Action  RequestAction  `json:"action" validate:"required,min=3,max=20"`
	Message string         `json:"message" validate:"required,min=3,max=50"`
	Changes map[string]any `json:"changes"`
}
