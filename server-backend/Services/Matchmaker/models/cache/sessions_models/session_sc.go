package sessions_models

type UserStatus string

const (
	StatusAFK    UserStatus = "afk"     // User is afk
	StatusOnline UserStatus = "online"  // User is online
	StatusInGame UserStatus = "in_game" // User is in a game
	StatusInRoom UserStatus = "in_room" // User is in a waiting room
)

type User struct {
	ID         string     `json:"id"`
	Username   string     `json:"username"`
	Status     UserStatus `json:"status"`
	InstanceID string     `json:"inst_id"`
}
