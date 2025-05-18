package cache_rooms

import (
	"slices"
	"time"

	"github.com/Math-Vov13/BloodyMoon/models/rooms_models"
	"github.com/Math-Vov13/BloodyMoon/models/users_models"
	"github.com/Math-Vov13/BloodyMoon/pkg/generator"
)

func CreateRoom(host *users_models.User, configs *rooms_models.RoomConfig) *rooms_models.RoomCreated {
	// Check if the user is not already in a room
	if _, exists := fake_cache[host.ID]; exists {
		return nil
	}

	// Generate a unique ID for the room
	id_generated := generator.GenerateID(32)
	for {
		if GetRoomByID(id_generated) == nil {
			break
		}
		id_generated = generator.GenerateID(32)
	}
	// Generate a unique code for the room
	code_generated := generator.GenerateID(6)
	for {
		if GetRoomWithCode(code_generated) == nil {
			break
		}
		code_generated = generator.GenerateID(6)
	}

	// Create a new game
	room := rooms_models.RoomCreated{
		RoomID:    id_generated,
		JoinCode:  code_generated,
		HostID:    host.ID,
		Status:    rooms_models.StatusDefault,
		Players:   []string{host.ID},
		CreatedAt: time.Now().Unix(),
		RoomConfig: rooms_models.RoomConfig{
			RoomName:   "room_name",
			MaxPlayers: 10,
			GameMode:   "default",
		},
	}

	// Store the game in Redis
	// err := rdb.Set("game:"+room.RoomID, room, 0).Err()
	// if err != nil {
	// 	return nil
	// }

	fake_cache[host.ID] = &room
	return &room
}

func JoinRoom(hostId string, roomId string) bool {
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

	if room.Status != rooms_models.StatusActive {
		return false
	}

	room.Players = append(room.Players, hostId)
	if len(room.Players) > room.RoomConfig.MaxPlayers {
		room.Status = rooms_models.StatusFull
	}

	return true
}

func LeaveRoom(hostId string, roomId string) bool {
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

	for _, room := range fake_cache {
		if room.RoomID == roomId {
			for i, player := range room.Players {
				if player == hostId {
					room.Players = slices.Delete(room.Players, i, i+1)
					return true
				}
			}
		}
	}

	return false
}

func DeleteRoom(host *users_models.User) {
	// Delete the game from Redis
	// err := rdb.Del("game:"+host.ID).Err()
	// if err != nil {
	// 	return
	// }

	delete(fake_cache, host.ID)
}
