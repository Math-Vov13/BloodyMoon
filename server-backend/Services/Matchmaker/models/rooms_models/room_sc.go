package rooms_models

type StatusType string

const (
	StatusDefault StatusType = "offline" // just created
	StatusActive  StatusType = "active"  // Host joined
	StatusFull    StatusType = "full"    // Room is full
)

type RoomConfig struct {
	RoomName   string `json:"room_name"`
	GameMode   string `json:"game_mode" default:"default"`
	MaxPlayers int    `json:"max_players" default:"10"`
	Landscape  string `json:"landscape" default:"default"`
}

type RoomCreated struct {
	RoomID    string     `json:"room_id"`
	JoinCode  string     `json:"join_code"`
	HostID    string     `json:"host"`
	Status    StatusType `json:"status" default:"offline"`
	Players   []string   `json:"players"`
	CreatedAt int64      `json:"created_at"`
	RoomConfig
}
