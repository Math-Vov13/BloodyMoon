package websocket_Room

import (
	"github.com/Math-Vov13/BloodyMoon/internal/database/cache_redis/cache_rooms"
	"github.com/Math-Vov13/BloodyMoon/internal/database/mongodb"
	"github.com/Math-Vov13/BloodyMoon/models/rooms_models"
	"github.com/Math-Vov13/BloodyMoon/models/users_models"
	"github.com/gin-gonic/gin"
)

func verifyAccesstoRoom(c *gin.Context, player *users_models.User) (room *rooms_models.RoomCreated, err string, code int) {
	room_code := c.Query("code")
	// 1. Check if the code is provided and valid
	if len(room_code) != 6 {
		err = "Invalid code"
		code = 400
		return
	}

	// 2. Check if the room exists
	room = cache_rooms.GetRoomWithCode(room_code)
	if room == nil {
		err = "Room not found"
		code = 404
		return
	}

	// 3. Check if the room is active
	if (room.Status != rooms_models.StatusActive) && room.HostID != player.ID {
		err = "Room is not active"
		code = 409
		return
	}

	// 4. Check if the room is full
	if len(room.Players) >= room.RoomConfig.MaxPlayers {
		err = "Room is full"
		code = 409
		return
	}

	// 5. Check if the player is already in a room / game
	if mongodb.IsInGame(player.ID) {
		err = "You are already in a game"
		code = 409
		return
	}

	// if slices.Contains(room.Players, player.ID) {
	// 	err = "You are already in the room"
	// 	code = 409
	// 	return
	// }

	return
}

func removeClientFromRoom(client *Client) {
	// // Remove the client from the room
	// mutex.Lock()
	// defer mutex.Unlock()

	// room := cache_rooms.GetRoomByID(client.roomID)
	// if room == nil {
	// 	return
	// }

	// // Remove the client from the room
	// for i, player := range room.Players {
	// 	if player == client.user.ID {
	// 		room.Players = append(room.Players[:i], room.Players[i+1:]...)
	// 		break
	// 	}
	// }

	// // Update the room in the cache
	// cache_rooms.ChangeRoomStatus(room.RoomID, rooms_models.StatusActive)
}
