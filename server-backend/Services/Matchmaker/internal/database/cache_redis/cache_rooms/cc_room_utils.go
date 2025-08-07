package cache_rooms

import (
	"fmt"

	"github.com/Math-Vov13/BloodyMoon/models/cache/rooms_models"
)

func GetRoomByID(room_id string) (*rooms_models.Room, error) {
	for _, room := range fake_cache {
		if room.ID == room_id {
			return room, nil
		}
	}
	return nil, fmt.Errorf("room not found")
}

func GetRoomByCode(code string) (*rooms_models.Room, error) {
	// Get the game from Redis
	// val, err := rdb.Get("game:" + code).Result()
	// if err != nil {
	// 	return nil
	// }

	// var room rooms_models.RoomCreated
	// err = json.Unmarshal([]byte(val), &room)
	// if err != nil {
	// 	return nil
	// }

	for _, room := range fake_cache {
		if room.JoinCode == code {
			return room, nil
		}
	}

	return nil, fmt.Errorf("room not found")
}
