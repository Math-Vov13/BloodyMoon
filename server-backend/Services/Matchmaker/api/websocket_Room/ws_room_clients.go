package websocket_Room

import (
	"fmt"

	"github.com/Math-Vov13/BloodyMoon/models/users_models"
	"github.com/gorilla/websocket"
)

// type Client struct {
// 	conn *websocket.Conn
// 	send chan []byte
// }

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

func (r *Room) ConnectBroadcast() {
	for {
		select {
		case message := <-r.broadcast:
			fmt.Println("Broadcasting message to clients:", message)
			for _, client := range r.clients {
				err := client.conn.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					// Gérer la déconnexion du client ici
				}
			}

		default: // If the channel is full, close the connection
			break
		}
	}
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
	delete(r.clients, client.user.ID)
	close(client.send)
	client.conn.Close()
}

func (r *Room) Broadcast(message []byte) {
	for _, client := range r.clients {
		select {
		case client.send <- message:
		default:
			r.RemoveClient(client)
		}
	}
}
