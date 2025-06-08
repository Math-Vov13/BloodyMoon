package websocket_Room

import (
	"fmt"
	"time"

	"github.com/Math-Vov13/BloodyMoon/models/users_models"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	send     chan []byte
	user     *users_models.User
	lastPing int64 // Timestamp of the last ping
}

type Room struct {
	ID        string
	clients   map[string]*Client
	broadcast chan []byte
	hostID    string
}

// --- Rooms ---
func CreateRoom(roomId string, hostId string) (room *Room) {
	room = &Room{
		ID:        roomId,
		clients:   make(map[string]*Client),
		broadcast: make(chan []byte),
		hostID:    hostId,
	}
	game_rooms[roomId] = room

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
func (r *Room) CreateClient(conn *websocket.Conn, player *users_models.User) (client *Client) {
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

func (r *Room) MutlicastMessage(message []byte, clients map[*Client]bool) {
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
