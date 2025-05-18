package users_models

type UserStatus string

const (
	StatusOffline UserStatus = "offline" // User is offline
	StatusOnline  UserStatus = "online"  // User is online
	StatusInGame  UserStatus = "in_game" // User in a game
	StatusInRoom  UserStatus = "in_room" // User in a waiting room
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	HOST     bool   `json:"host"`
}

type DBUser struct {
	ID       string     `json:"id"`
	Username string     `json:"username"`
	Status   UserStatus `json:"in_game"`
	GameID   string     `json:"game_id"`
}
