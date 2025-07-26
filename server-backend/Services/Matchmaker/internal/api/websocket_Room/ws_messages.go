package websocket_Room

import (
	"encoding/json"
	"fmt"

	"github.com/Math-Vov13/BloodyMoon/models/requests_models"
	"github.com/Math-Vov13/BloodyMoon/models/responses_models"
)

func prepareMessage(message any) (msg []byte) {
	// Prepare the message to be sent to the client
	msg, err := json.Marshal(message)
	if err != nil {
		msg, _ = json.Marshal(responses_models.ResponseForError{
			BaseResponse: responses_models.BaseResponse{
				Code:    500,
				Type:    responses_models.TypeError,
				Message: "Error when sending message",
			},
			Error: err.Error(),
		})
		return
	}
	return
}

func decodeMessage(p []byte) (msg *requests_models.RequestEvent, code int, err error) {
	code = 200 // Default success code

	// Verify message length
	if len(p) == 0 {
		code = 400 // Bad Request
		err = fmt.Errorf("empty message received")
		return
	}

	// Unmarshal JSON into the Message struct
	if err1 := json.Unmarshal(p, &msg); err1 != nil {
		code = 400 // Bad Request
		err = fmt.Errorf("error decoding message: %w", err1)
		return
	}

	// Validate the decoded message
	if err1 := validate.Struct(msg); err1 != nil {
		code = 422 // Unprocessable Entity
		err = fmt.Errorf("validation error: %w", err1)
		return
	}

	// Check if the message type is valid
	if msg.Action != requests_models.TypeMessage && msg.Action != requests_models.TypeConfig && msg.Action != requests_models.TypePlay {
		code = 422 // Unprocessable Entity
		err = fmt.Errorf("invalid request action: %v", msg.Action)
		return
	}

	return
}
