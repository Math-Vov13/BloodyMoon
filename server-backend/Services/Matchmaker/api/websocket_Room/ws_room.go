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

const MAX_CLIENTS = 100000

var (
	game_rooms   = make(map[string]*Room)
	conn_clients = make(map[*Client]bool)
	mutex        = sync.Mutex{}
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

	// Verify WebSocket maximum number of clients
	if len(conn_clients) >= MAX_CLIENTS {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"success": false,
			"message": "WebSocket is overloaded, please try again later",
			"code":    503,
			"type":    "error",
		})
		return
	}

	// Get the room
	actual_room := game_rooms[room_target.RoomID]
	if actual_room == nil {
		actual_room = &Room{
			ID:        room_target.RoomID,
			clients:   make(map[string]*Client),
			broadcast: make(chan []byte),
		}
		game_rooms[room_target.RoomID] = actual_room
		go actual_room.ConnectBroadcast() // Start the broadcast goroutine
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
		return
	}

	fmt.Printf("Vous êtes dans la room: %s\n", room_target.RoomID)
	fmt.Printf("Vous êtes host: %t\n", is_host)

	// Create the client
	fmt.Printf("(nbr: %d) ", len(conn_clients))
	fmt.Println(">> Client connecté :", conn.RemoteAddr())
	mutex.Lock()
	client := actual_room.CreateClient(conn, user) // Create client and add it to the room
	conn_clients[client] = true                    // Add the client to the map
	mutex.Unlock()

	conn.WriteJSON(gin.H{
		"success": true,
		"message": "You are now connected to the room!",
		"room":    room_target,
		"owner":   is_host,
		"type":    "info",
	})

	// go func() {
	// 	for msg := range client.send {
	// 		fmt.Printf("Envoi du message à %s: %s\n", conn.RemoteAddr(), msg)
	// 		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
	// 			break
	// 		}
	// 	}
	// }()

	// --- Read messages from the client ---
	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		switch mt {
		case websocket.TextMessage:
			fmt.Printf("Message reçu de %s: %b\n", conn.RemoteAddr(), msg)
			actual_room.broadcast <- msg // Broadcast the message to all clients in the room
			//broadcast <- Message{client: client, content: msg}             // Save to broadcast channel

		case websocket.PongMessage:
			// Optionnel : log du pong
			log.Println("Pong reçu")

		}
	}

	// Déconnexion
	mutex.Lock()
	actual_room.RemoveClient(client) // Remove the client from the room
	delete(conn_clients, client)     // Remove the client from the map
	mutex.Unlock()

	fmt.Println(">> Client déconnecté :", conn.RemoteAddr())
}
