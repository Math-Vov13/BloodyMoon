package endpoints

import (
	"fmt"
	"net/http"

	"github.com/Math-Vov13/BloodyMoon/internal/database/cache_redis/cache_rooms"
	"github.com/Math-Vov13/BloodyMoon/internal/database/mongodb"
	"github.com/Math-Vov13/BloodyMoon/models/rooms_models"
	"github.com/Math-Vov13/BloodyMoon/models/users_models"
	"github.com/gin-gonic/gin"
)

func CreatePrivateGame(c *gin.Context) {
	var user *users_models.User = c.MustGet("user").(*users_models.User)

	if mongodb.IsInGame(user.ID) {
		// TODO:  Vérifie si la partie est terminée ou non
		c.JSON(http.StatusConflict, gin.H{
			"message": "You are already in a room",
		})
		return
	}

	room := cache_rooms.CreateRoom(user, &rooms_models.RoomConfig{
		RoomName: "room_name",
		GameMode: "default",
	})
	if room == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "You can't create a room",
		})
		return
	}

	fmt.Println("Room created: ", room)

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Your private room has been created",
		"room_id":   room.RoomID,
		"room_code": room.JoinCode,
		"status":    room.Status,
		"config":    room.RoomConfig,
	})
}

func DeletePrivateGame(c *gin.Context) {
	// Vérifie si le joueur peut supprimer une partie privée
	// 1. Il est le créateur de la partie
	// 2. La partie n'est pas encore commencée

	// Supprimer la partie

	c.JSON(http.StatusOK, gin.H{
		"message": "Your private room has been deleted",
		"room_id": "123456789",
	})
}
