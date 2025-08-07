package cache_rooms

import (
	"fmt"
	"time"

	"github.com/Math-Vov13/BloodyMoon/models/cache/rooms_models"
	"github.com/Math-Vov13/BloodyMoon/pkg/generator"
)

func CreateRoom(User_id string, configs *rooms_models.RoomConfig) (*rooms_models.Room, error) {
	// Check if the user is not already in a room
	if _, exists := fake_cache[User_id]; exists {
		return nil, fmt.Errorf("user already in a room")
	}

	// Generate a unique ID for the room
	id_generated := generator.GenerateID(32)
	for {
		if _, err := GetRoomByID(id_generated); err != nil {
			break
		}
		id_generated = generator.GenerateID(32)
	}
	// Generate a unique code for the room
	code_generated := generator.GenerateID(6)
	for {
		if _, err := GetRoomByCode(code_generated); err != nil {
			break
		}
		code_generated = generator.GenerateID(6)
	}

	// Create a new game
	room := &rooms_models.Room{
		ID:        id_generated,
		JoinCode:  code_generated,
		HostID:    User_id,
		Status:    rooms_models.StatusDefault,
		Players:   []string{User_id},
		CreatedAt: time.Now().Unix(),
		RoomConfig: rooms_models.RoomConfig{
			RoomName:   "room_name",
			MaxPlayers: 16,
			GameMode:   "default",
		},
	}

	// Store the game in Redis
	// err := rdb.Set("game:"+room.RoomID, room, 0).Err()
	// if err != nil {
	// 	return nil
	// }

	fake_cache[User_id] = room
	return room, nil
}

func ChangeRoomStatus(roomId string, status rooms_models.StatusType) error {
	room, err := GetRoomByID(roomId)
	if err != nil {
		return err
	}
	room.Status = status
	return nil
}

func DeleteRoom(roomID string) error {
	// Delete the game from Redis
	// err := rdb.Del("game:"+host.ID).Err()
	// if err != nil {
	// 	return
	// }

	room, err := GetRoomByID(roomID)
	if err != nil {
		return err
	}

	delete(fake_cache, room.HostID)
	return nil
}

func regenerateRoom(room *rooms_models.Room) (*rooms_models.Room, error) {
	fake_cache[room.HostID] = room
	return room, nil
}
