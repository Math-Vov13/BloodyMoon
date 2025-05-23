package websocket_Room

import (
	"fmt"

	"github.com/Math-Vov13/BloodyMoon/models/users_models"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
	user *users_models.User
}

type Room struct {
	ID        string
	clients   map[string]*Client
	broadcast chan []byte
}

// --- Rooms ---
func CreateRoom(roomId string) (room *Room) {
	room = &Room{
		ID:        roomId,
		clients:   make(map[string]*Client),
		broadcast: make(chan []byte),
	}
	game_rooms[roomId] = room

	return
}

func (r *Room) RemoveRoom() {
	// Remove the room from the map
	delete(game_rooms, r.ID)
	// Close the broadcast channel
	close(r.broadcast)
	// Close all client connections
	for _, client := range r.clients {
		client.conn.Close()
	}
}

// --- Clients ---
func (r *Room) CreateClient(conn *websocket.Conn, player *users_models.User) (client *Client) {
	// Create a new client
	client = &Client{
		conn: conn,
		send: make(chan []byte),
		user: player,
	}
	// Add the client to the room
	r.clients[player.ID] = client

	return
}

func (r *Room) RemoveClient(client *Client) {
	fmt.Printf("Client %s kicked from room %s\n", client.user.ID, r.ID)
	delete(r.clients, client.user.ID)
	close(client.send)
	client.conn.Close()
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
			r.RemoveClient(client)
		}
	}
}

func (r *Room) MutlicastMessage(message []byte, clients map[*Client]bool) {
	for client := range clients {
		select {
		case client.send <- message:
		default:
			r.RemoveClient(client)
		}
	}
}
