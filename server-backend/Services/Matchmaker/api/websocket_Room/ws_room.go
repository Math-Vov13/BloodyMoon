package websocket_Room

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Math-Vov13/BloodyMoon/internal/database/cache_redis/cache_rooms"
	"github.com/Math-Vov13/BloodyMoon/internal/database/mongodb"
	"github.com/Math-Vov13/BloodyMoon/models/rooms_models"
	"github.com/Math-Vov13/BloodyMoon/models/users_models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

type Message struct {
	client  *Client
	content []byte
}

var (
	all_clients = make(map[*Client]bool)
	broadcast   = make(chan Message)
	mutex       = sync.Mutex{}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // autorise toutes les origines
	},
}

func HandleWebSocket(c *gin.Context) {
	user := c.MustGet("user").(*users_models.User)

	// --- Verify Room and Player Access ---
	room_target, err1, errcode := verifyAccesstoRoom(c, user)
	if err1 != "" {
		c.JSON(errcode, gin.H{
			"success": false,
			"message": "Error connecting to the room: " + err1,
			"code":    errcode,
			"type":    "error",
		})
		return
	}

	// --- Upgrade the connection to a WebSocket ---
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Erreur WebSocket :", err)
		return
	}
	defer conn.Close()

	// --- Add the player to the room ---
	var is_host = (user.ID == room_target.HostID)
	if is_host {
		cache_rooms.ChangeRoomStatus(room_target.RoomID, rooms_models.StatusActive)
	} else {
		// TODO : CAUTION ! This solution does not work for AFK players
		cache_rooms.AddPlayerToRoom(user.ID, room_target.RoomID)
	}

	// Change the user's status to "in room"
	if succ := mongodb.ChangeUserStatus(user.ID, users_models.StatusInRoom, room_target.RoomID); !succ {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error connecting to the room: " + "Internal error",
			"code":    500,
			"type":    "error",
		})
	}

	fmt.Printf("Vous êtes dans la room: %s\n", room_target.RoomID)
	fmt.Printf("Vous êtes host: %t\n", is_host)

	// New Client connected
	fmt.Printf("(nbr: %d) ", len(all_clients))
	fmt.Println(">> Client connecté :", conn.RemoteAddr())
	client := &Client{conn: conn, send: make(chan []byte)} // Create Client instance

	go func() {
		for msg := range client.send {
			fmt.Printf("Envoi du message à %s: %s\n", conn.RemoteAddr(), msg)
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				break
			}
		}
	}()

	mutex.Lock()
	all_clients[client] = true // Register the client
	mutex.Unlock()

	conn.WriteJSON(gin.H{
		"success": true,
		"message": "You are now connected to the room!",
		"room":    room_target,
		"host":    is_host,
		"type":    "info",
	})

	// Read messages from the client
	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		switch mt {
		case websocket.TextMessage:
			fmt.Printf("Message reçu de %s: %s\n", conn.RemoteAddr(), msg)
			broadcast <- Message{client: client, content: msg} // Save to broadcast channel

		case websocket.PongMessage:
			// Optionnel : log du pong
			log.Println("Pong reçu")

		}
	}

	// Déconnexion
	mutex.Lock()
	delete(all_clients, client)
	mutex.Unlock()
	close(client.send)
	fmt.Println(">> Client déconnecté :", conn.RemoteAddr())
}

func StartMessageDispatcher() {
	go func() { // Start a goroutine to handle broadcasting messages
		for {
			msg := <-broadcast // Wait for a message to be sent

			mutex.Lock()
			for c := range all_clients {
				select {
				case c.send <- msg.content: // Send the message to the channel client
				default: // If the channel is full, close the connection
					fmt.Printf("Fermeture de la connexion avec %s\n", c.conn.RemoteAddr())
					close(c.send)
					delete(all_clients, c)
				}
			}
			mutex.Unlock()
		}
	}()
}
