package cache_sessions

import (
	"fmt"

	"github.com/Math-Vov13/BloodyMoon/models/cache/sessions_models"
)

func GetUserByName(username string) (*sessions_models.User, error) {
	// Tests purposes only!
	for _, user := range fake_cache {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func GetUserById(user_id string) (*sessions_models.User, error) {
	user, ok := fake_cache[user_id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func ChangeUserStatus(User_id string, Room_Id string, Status sessions_models.UserStatus) error {
	if _, ok := fake_cache[User_id]; ok {
		fake_cache[User_id].InstanceID = Room_Id
		fake_cache[User_id].Status = Status
		return nil
	} else {
		return fmt.Errorf("user not found")
	}
}

func GetUserRoomId(User_id string) (string, error) {
	user, err := GetUserById(User_id)
	if err != nil {
		return "", err
	}
	return user.InstanceID, nil
}
