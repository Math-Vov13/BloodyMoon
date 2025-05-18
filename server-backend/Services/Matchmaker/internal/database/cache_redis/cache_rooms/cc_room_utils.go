package cache_rooms

import "github.com/Math-Vov13/BloodyMoon/models/rooms_models"

func GetRoomByID(room_id string) *rooms_models.RoomCreated {
	for _, room := range fake_cache {
		if room.RoomID == room_id {
			return room
		}
	}

	return nil
}

func GetRoomWithCode(code string) *rooms_models.RoomCreated {
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
			return room
		}
	}

	return nil
}
