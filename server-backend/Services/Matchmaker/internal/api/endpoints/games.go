package endpoints

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Math-Vov13/BloodyMoon/internal/database/cache_redis/cache_rooms"
	"github.com/Math-Vov13/BloodyMoon/internal/database/cache_redis/cache_sessions"
	"github.com/Math-Vov13/BloodyMoon/models/cache/rooms_models"
	"github.com/Math-Vov13/BloodyMoon/models/cache/sessions_models"
	"github.com/gin-gonic/gin"
)

func CreatePrivateGame(c *gin.Context) {
	var user *sessions_models.User = c.MustGet("user").(*sessions_models.User)
	c.Writer.Header().Set("Server", "localhost")
	c.Writer.Header().Set("Origin", os.Getenv("URL"))
	c.Writer.Header().Set("Cache-Control", "no-store") // no cache

	// Check if the user is already in a room
	if user.InstanceID != "" {
		// TODO:  Vérifie si la partie est terminée ou non
		//c.Writer.Header().Set("Cache-Control", "max-age=120") // no cache
		c.JSON(http.StatusConflict, gin.H{
			"message": "You are already in a room",
			"error":   "You are already in a room or game",
			"room_id": user.InstanceID,
		})
		return
	}

	// Check if the user has already created a room (but not joined it yet)
	if _, err := cache_rooms.GetRoomForHostID(user.ID); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "You are already in a room",
			"error":   err.Error(),
			"room_id": user.InstanceID,
		})
		return
	}

	// Create a Room
	room, err := cache_rooms.CreateRoom(user.ID, &rooms_models.RoomConfig{
		RoomName: "room_name",
		GameMode: "default",
	})
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Error creating room",
			"error":   err.Error(),
		})
		return
	}

	// Change Player Status
	// if cache_sessions.ChangeUserStatus(user.ID, room.ID, sessions_models.StatusOnline) != nil {
	// 	c.JSON(http.StatusConflict, gin.H{
	// 		"message": "Error with your profile status?",
	// 	})
	// 	return
	// }

	fmt.Println("Room created: ", room)
	c.JSON(http.StatusCreated, gin.H{
		"message":   "Your private room has been created",
		"room_id":   room.ID,
		"room_code": room.JoinCode,
		"status":    room.Status,
		"config":    room.RoomConfig,
	})
}

func DeletePrivateGame(c *gin.Context) {
	var user *sessions_models.User = c.MustGet("user").(*sessions_models.User)
	c.Writer.Header().Set("Server", "localhost")
	c.Writer.Header().Set("Origin", os.Getenv("URL"))
	c.Writer.Header().Set("Cache-Control", "no-store") // no cache

	// Vérifie si le joueur peut supprimer une partie privée
	// 1. Il est le créateur de la partie
	// 2. La partie n'est pas encore commencée

	if user.InstanceID == "" {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"message": "User is not connected to a room",
		})
		return
	}

	room, err := cache_rooms.GetRoomByID(user.InstanceID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Room is expired!",
		})
		return
	}

	if room.HostID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "You have no rights to delete this room !",
			"error": gin.H{
				"needed":    "HOST",
				"privilege": "MEMBER",
			},
		})
		return
	}

	if len(room.Players) > 1 {
		// Déconnecte tous les joueurs de la partie ? (connexion au serveur websocket)
		c.JSON(http.StatusConflict, gin.H{
			"message": "Room is not umpty!",
		})
		return
	}

	// Supprimer la partie
	if err := cache_rooms.DeleteRoom(room.ID); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Room can't be deleted!",
			"error":   err.Error(),
		})
		return
	}

	// Changer le status du Joueur
	if cache_sessions.ChangeUserStatus(user.ID, "", sessions_models.StatusOnline) != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Error with your profile status?",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your private room has been deleted",
		"room_id": room.ID,
	})
}
