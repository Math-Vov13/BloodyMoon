package cache_rooms

import (
	"slices"

	"github.com/Math-Vov13/BloodyMoon/models/rooms_models"
)

func ChangeRoomStatus(roomdId string, status rooms_models.StatusType) bool {
	room := GetRoomByID(roomdId)
	room.Status = status
	return true
}

func AddPlayerToRoom(playerId string, roomId string) bool {
	// Get the game from Redis
	// val, err := rdb.Get("game:" + roomId).Result()
	// if err != nil {
	// 	return false
	// }

	// var room rooms_models.RoomCreated
	// err = json.Unmarshal([]byte(val), &room)
	// if err != nil {
	// 	return false
	// }

	room := GetRoomByID(roomId)
	if room == nil {
		return false
	}

	room.Players = append(room.Players, playerId)
	if len(room.Players) >= room.RoomConfig.MaxPlayers {
		room.Status = rooms_models.StatusFull
	}

	return true
}

func RemovePlayerFromRoom(playerId string, roomId string) bool {
	// Get the game from Redis
	// val, err := rdb.Get("game:" + roomId).Result()
	// if err != nil {
	// 	return false
	// }

	// var room rooms_models.RoomCreated
	// err = json.Unmarshal([]byte(val), &room)
	// if err != nil {
	// 	return false
	// }

	room := GetRoomByID(roomId)
	if room == nil {
		return false
	}

	for i, player := range room.Players {
		if player == playerId {
			room.Players = slices.Delete(room.Players, i, i+1)
			break
		}
	}

	if len(room.Players) < room.RoomConfig.MaxPlayers {
		room.Status = rooms_models.StatusActive
	}

	return true
}
