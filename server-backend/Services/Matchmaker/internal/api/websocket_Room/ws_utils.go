package websocket_Room

import (
	"fmt"

	"github.com/Math-Vov13/BloodyMoon/internal/database/cache_redis/cache_rooms"
	"github.com/Math-Vov13/BloodyMoon/models/cache/rooms_models"
	"github.com/Math-Vov13/BloodyMoon/models/cache/sessions_models"
	"github.com/gin-gonic/gin"
)

func verifyAccesstoRoom(c *gin.Context, player *sessions_models.User) (room *rooms_models.Room, code int, err error) {
	room_code := c.Query("code")
	// 1. Check if the code is provided and valid
	if len(room_code) != 6 {
		err = fmt.Errorf("invalid code")
		code = 400
		return
	}

	// 2. Check if the room exists
	room, err = cache_rooms.GetRoomByCode(room_code)
	if err != nil {
		code = 404
		return
	}

	// 3. Check if the room is active
	if (room.Status != rooms_models.StatusActive) && room.HostID != player.ID {
		err = fmt.Errorf("Room is not active")
		code = 409
		return
	}

	// 4. Check if the room is full
	if len(room.Players) >= room.RoomConfig.MaxPlayers {
		err = fmt.Errorf("Room is full")
		code = 409
		return
	}

	// 5. Check if the player is already in a room / game
	if player.InstanceID != "" {
		err = fmt.Errorf("you are already in a room")
		code = 409
		return
	}

	// 6. Check if the player is banned from the room
	// if room.IsBanned(player.ID) {
	// 	err = fmt.Errorf("you are banned from this room")
	// 	code = 403
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
