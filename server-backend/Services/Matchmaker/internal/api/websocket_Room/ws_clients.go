package websocket_Room

import (
	"fmt"
	"time"

	"github.com/Math-Vov13/BloodyMoon/models/cache/sessions_models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	send     chan []byte
	user     *sessions_models.User
	lastPing int64 // Timestamp of the last ping
}

type Room struct {
	ID        string
	clients   map[string]*Client
	broadcast chan []byte
	hostID    string
	uuid      uuid.UUID
}

// --- Rooms ---
func CreateRoom(roomId string, hostId string) (room *Room) {
	room = &Room{
		ID:        roomId,
		clients:   make(map[string]*Client),
		broadcast: make(chan []byte),
		hostID:    hostId,
		uuid:      uuid.NewSHA1(uuid.NameSpaceURL, []byte(roomId)),
	}
	game_rooms[roomId] = room

	return
}

func (r *Room) ChangeRoomHost() (newHost *sessions_models.User, err error) {
	if r == nil {
		fmt.Println("Tried to change host of nil Room")
		err = fmt.Errorf("nil Room")
		return
	}
	if len(r.clients) == 0 {
		fmt.Println("No clients in room to change host")
		err = fmt.Errorf("no clients in room %s", r.ID)
		return
	}

	for id, player := range r.clients {
		if id != r.hostID {
			r.hostID = id // Change the host to the first client found
			newHost = player.user
			fmt.Printf("New host for room %s is %s\n", r.ID, newHost.ID)
			return
		}
	}

	err = fmt.Errorf("no suitable new host found for room %s", r.ID)
	return
}

func (r *Room) RemoveRoom() (err error) {
	if r == nil {
		fmt.Println("Tried to remove nil Room")
		err = fmt.Errorf("nil Room")
		return
	}
	// Remove the room from the map
	delete(game_rooms, r.ID)
	// Close the broadcast channel
	close(r.broadcast)
	// Close all client connections
	for _, client := range r.clients {
		client.conn.Close()
	}
	return
}

// --- Clients ---
func (r *Room) CreateClient(conn *websocket.Conn, player *sessions_models.User) (client *Client) {
	// Create a new client
	client = &Client{
		conn:     conn,                  // WebSocket connection
		send:     make(chan []byte, 32), // Channel for sending messages to the client
		user:     player,                // User associated with the client
		lastPing: time.Now().Unix(),     // Initialize lastPing to 0
	}
	// Add the client to the room
	r.clients[player.ID] = client

	return
}

func (r *Room) RemoveClient(client *Client) (err error) {
	if r == nil || client == nil {
		fmt.Println("Tried to remove nil Client or Room")
		err = fmt.Errorf("nil Client or Room")
		return
	}
	if _, exists := r.clients[client.user.ID]; !exists {
		fmt.Printf("Tried to kick unexisting Client '%s' in room (%s)\n", client.user.ID, r.ID)
		err = fmt.Errorf("client %s not found in room %s", client.user.ID, r.ID)
		return
	}

	fmt.Printf("Client %s kicked from room %s\n", client.user.ID, r.ID)
	delete(r.clients, client.user.ID)
	close(client.send)
	client.conn.Close()
	return
}

// --- Messages ---
func (r *Room) GenerateUUID() string {
	return uuid.NewSHA1(r.uuid, fmt.Appendf(nil, "%d", time.Now().UnixNano())).String()
}

func (r *Room) BroadcastMessage(message []byte, excludeIds map[string]bool) {
	for id, client := range r.clients {
		if excludeIds != nil && excludeIds[id] {
			continue
		}
		select {
		case client.send <- message:
		default:
			mutex.Lock()
			client.conn.Close()
			mutex.Unlock()
		}
	}
}

// func (r *Room) SpecialMessage(main_message []byte, special_demands map[*Client][]byte) {
// 	for client, message := range special_demands {
// 		select {
// 		case client.send <- message:
// 		default:
// 			mutex.Lock()
// 			client.conn.Close()
// 			mutex.Unlock()
// 		}
// 	}
// 	if len(special_demands) == 0 {
// 		// If no special demands, send the main message to all clients
// 		r.BroadcastMessage(main_message, nil)
// 		return
// 	}
// }

func (r *Room) MulticastMessage(message []byte, clients map[*Client]bool) {
	for client := range clients {
		select {
		case <-client.send:
			continue // Skip if the client is already sending a message
		case client.send <- message:
		default:
			mutex.Lock()
			client.conn.Close()
			mutex.Unlock()
		}
	}
}
