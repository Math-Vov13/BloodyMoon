package rooms_models

type StatusType string

const (
	StatusDefault StatusType = "offline" // just created
	StatusActive  StatusType = "active"  // Host joined
	StatusFull    StatusType = "full"    // Room is full
)

type RoomConfig struct {
	RoomName   string `json:"room_name"`
	GameMode   string `json:"game_mode"`
	MaxPlayers int    `json:"max_players"`
	Landscape  string `json:"landscape"`
}

type RoomCreated struct {
	RoomID    string     `json:"room_id"`
	JoinCode  string     `json:"join_code"`
	HostID    string     `json:"host"`
	Status    StatusType `json:"status"`
	Players   []string   `json:"players"`
	CreatedAt int64      `json:"created_at"`
	RoomConfig
}
