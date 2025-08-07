package responses_models

import "github.com/gin-gonic/gin"

type ResponseType string
type ResponseDest string

const (
	TypeError   ResponseType = "Error"   // Error message
	TypeSystem  ResponseType = "System"  // System message > Player Joined, Player Left, Start Game
	TypeMessage ResponseType = "Message" // Message from chat
	TypeAck     ResponseType = "Ack"     // Acknowledgment message
	// TypeUpdate  ResponseType = "Update"  // Configuration message > Host changed Game config
)

const (
	DestinatorUser ResponseDest = "User" // User only
	DestinatorAll  ResponseDest = "All"  // All users in the room
)

type BaseResponse struct {
	// Code    int          `json:"code"`
	EventId string       `json:"id"`
	Type    ResponseType `json:"type"`
	Message string       `json:"message"`
}

type ResponseForMessage struct {
	BaseResponse
	Content string `json:"content"`
	Author  string `json:"author"`
}

type ResponseForError struct {
	BaseResponse
	Error string `json:"error"`
}

type ResponseForSystem struct {
	BaseResponse
	Destinator ResponseDest `json:"dest"`
	Content    gin.H        `json:"content"`
}

type ResponseForAck struct {
	BaseResponse
	AckId   string `json:"ack_id"`
	Content gin.H  `json:"content"`
}
