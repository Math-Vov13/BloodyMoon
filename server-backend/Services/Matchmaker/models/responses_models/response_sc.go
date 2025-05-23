package responses_models

import "github.com/gin-gonic/gin"

type ResponseType string

const (
	TypeError   ResponseType = "Error"   // Error message
	TypeSystem  ResponseType = "System"  // System message > Player Joined, Player Left, Start Game
	TypeMessage ResponseType = "Message" // Message from chat
	TypeUpdate  ResponseType = "Update"  // Configuration message > Host changed Game config
)

type BaseResponse struct {
	Code    int          `json:"code"`
	Type    ResponseType `json:"type"`
	Message string       `json:"message"`
}

type ResponseForMessage struct {
	BaseResponse
	Content  string `json:"content"`
	Username string `json:"username"`
}

type ResponseForError struct {
	BaseResponse
	Error string `json:"error"`
}

type ResponseForSystem struct {
	BaseResponse
	Content gin.H `json:"content"`
}
