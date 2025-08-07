package endpoints

import (
	"os"

	"github.com/Math-Vov13/BloodyMoon/models/cache/rooms_models"
	"github.com/Math-Vov13/BloodyMoon/models/cache/sessions_models"
	"github.com/gin-gonic/gin"
)

func ConnectToRoom(c *gin.Context) {
	c.Writer.Header().Set("Server", "localhost")
	c.Writer.Header().Set("Origin", os.Getenv("URL"))
	c.Writer.Header().Set("Cache-Control", "no-store") // no cache

	user := c.MustGet("user").(*sessions_models.User)
	room := c.MustGet("room").(*rooms_models.Room)

	if user.InstanceID != "" {
		c.JSON(409, gin.H{
			"error":   "You are already in a room or game",
			"room_id": user.InstanceID,
		})
		return
	}

	// Generate JWT for WebSocket Connection

	c.JSON(200, gin.H{
		"room_id":   room.ID,
		"room_code": room.JoinCode,
		"status":    room.Status,
		"config":    room.RoomConfig,
		"host_id":   room.HostID,
		"players":   room.Players,
	})
}

// func DisconnectFromRoom(c *gin.Context) {
// 	c.Writer.Header().Set("Server", "localhost")
// 	c.Writer.Header().Set("Origin", os.Getenv("URL"))
// 	c.Writer.Header().Set("Cache-Control", "no-store") // no cache

// 	user := c.MustGet("user").(*sessions_models.User)
// 	room := c.MustGet("room").(*rooms_models.Room)

// 	if user.InstanceID == "" {
// 		c.JSON(409, gin.H{
// 			"error": "You are not in a room or game",
// 		})
// 		return
// 	}

// 	// Logic to disconnect the user from the room
// 	user.InstanceID = ""
// 	room.RemovePlayer(user.ID)

// 	c.JSON(200, gin.H{
// 		"message": "Successfully disconnected from the room",
// 	})
// }
