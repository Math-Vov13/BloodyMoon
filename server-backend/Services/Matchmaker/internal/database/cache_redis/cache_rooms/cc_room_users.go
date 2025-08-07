package cache_rooms

import (
	"slices"

	"github.com/Math-Vov13/BloodyMoon/models/cache/rooms_models"
)

func AddPlayerToRoom(playerId string, roomId string) error {
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

	room, err := GetRoomByID(roomId)
	if err != nil {
		return err
	}

	room.Players = append(room.Players, playerId)
	if len(room.Players) >= room.RoomConfig.MaxPlayers {
		room.Status = rooms_models.StatusFull
	}

	return nil
}

func RemovePlayerFromRoom(playerId string, roomId string) error {
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

	room, err := GetRoomByID(roomId)
	if err != nil {
		return err
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

	return nil
}

func ChangeRoomHost(roomId string, newOwnerId string) error {
	room, err := GetRoomByID(roomId)
	if err != nil {
		return err
	}

	DeleteRoom(roomId) // Remove the old room from cache
	room.HostID = newOwnerId
	regenerateRoom(room) // Create a new room with the updated owner
	return nil
}

func GetRoomForHostID(hostID string) (*rooms_models.Room, error) {
	for _, room := range fake_cache {
		if room.HostID == hostID {
			return room, nil
		}
	}
	return nil, nil // No room found for the given host ID
}
