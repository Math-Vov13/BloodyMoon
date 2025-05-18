package mongodb

import "github.com/Math-Vov13/BloodyMoon/models/users_models"

var fake_db = map[string]*users_models.DBUser{
	"1234": {
		ID:       "1234",
		Username: "test1",
		Status:   "offline",
	},
	"5678": {
		ID:       "5678",
		Username: "test2",
		Status:   "offline",
	},
}

// func CreateUser(username string) *users_models.DBUser {
// 	// Create a new user in the database
// 	// err := rdb.Set("user:"+userID, user, 0).Err()
// 	// if err != nil {
// 	// 	return nil
// 	// }
// 	user := &users_models.DBUser{
// 		ID:       "user_id",
// 		Username: username,
// 		Status:   "offline",
// 	}
// 	fake_db[user.ID] = user
// 	return user
// }

func GetUserByName(username string) *users_models.DBUser {
	for _, user := range fake_db {
		if user.Username == username {
			return user
		}
	}
	return nil
}

func ChangeUserStatus(userID string, status users_models.UserStatus, gameID string) bool {
	// Change the user's status in the database
	// err := rdb.Set("user:"+userID+":in_game", inGame, 0).Err()
	// if err != nil {
	// 	return
	// }
	if _, ok := fake_db[userID]; ok {
		fake_db[userID].GameID = gameID
		fake_db[userID].Status = status
		return true
	} else {
		return false
	}
}

func IsInGame(userID string) bool {
	// Check if the user is in a game
	// err := rdb.Get("user:" + userID + ":in_game").Err()
	// if err != nil {
	// 	return false
	// }
	// return true
	if _, ok := fake_db[userID]; ok {
		return fake_db[userID].Status == users_models.StatusInRoom || fake_db[userID].Status == users_models.StatusInGame
	}
	return false
}
