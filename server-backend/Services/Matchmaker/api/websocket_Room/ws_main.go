package websocket_Room

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/Math-Vov13/BloodyMoon/internal/database/cache_redis/cache_rooms"
	"github.com/Math-Vov13/BloodyMoon/internal/database/mongodb"
	"github.com/Math-Vov13/BloodyMoon/models/requests_models"
	"github.com/Math-Vov13/BloodyMoon/models/responses_models"
	"github.com/Math-Vov13/BloodyMoon/models/rooms_models"
	"github.com/Math-Vov13/BloodyMoon/models/users_models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
)

const MAX_CLIENTS = 100000

// Validator instance
var validate = validator.New()

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
		c.Writer.Header().Set("Retry-After", "15")         // 15 seconds
		c.Writer.Header().Set("Cache-Control", "no-cache") // no cache
		c.Writer.Header().Set("Connection", "close")       // close connection
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"success": false,
			"message": "WebSocket is overloaded, please try again later",
			"code":    503,
			"type":    "error",
		})
		return
	}

	// TODO : Upgrade la connexion dès le début
	// --- Upgrade the connection to a WebSocket ---
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusUpgradeRequired, gin.H{
			"success": false,
			"message": "Please use a WebSocket connection",
			"code":    http.StatusUpgradeRequired,
			"type":    "error",
		})
		log.Println("Erreur WebSocket :", err)
		return
	}
	defer conn.Close()

	// --- Add the player to the room ---
	// Get the room
	actual_room := game_rooms[room_target.RoomID]
	if actual_room == nil {
		actual_room = CreateRoom(room_target.RoomID) // Create the room if it doesn't exist
	}

	// Check if the player is already in the room
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
	mutex.Lock()
	client := actual_room.CreateClient(conn, user) // Create client and add it to the room
	conn_clients[client] = true                    // Add the client to the map
	mutex.Unlock()
	fmt.Printf(">> Client connecté : %s (total: %d)\n", conn.RemoteAddr(), len(conn_clients))

	go func() {
		for msg := range client.send {
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				break
			}
		}
	}()

	// Say Welcome to the client !
	// --> client
	client.send <- prepareMessage(gin.H{
		"success": true,
		"message": "You are now connected to the room!",
		"room":    room_target,
		"owner":   is_host,
		"type":    "info",
	}) // Send the message to the client

	// --> all clients in the room
	actual_room.BroadcastMessage(prepareMessage(responses_models.ResponseForSystem{
		BaseResponse: responses_models.BaseResponse{
			Code:    200,
			Type:    responses_models.TypeSystem,
			Message: fmt.Sprintf("User '%s' has joined the room", user.Username),
		},
		Content: gin.H{
			"status":   "joined",
			"user_id":  user.ID,
			"username": user.Username,
		},
	}), map[string]bool{
		user.ID: true,
	})

	// --- Read messages from the client ---
	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		switch mt {
		case websocket.TextMessage:
			// Unmarshal JSON into the Message struct
			var MessageRequest requests_models.RequestEvent
			if err := json.Unmarshal(msg, &MessageRequest); err != nil {
				log.Println("JSON unmarshal error:", err)
				client.send <- prepareMessage(gin.H{
					"success": false,
					"message": "Error decoding message: " + err.Error(),
					"code":    400,
					"error":   err.Error(),
					"type":    "error",
				})
				continue
			}

			// Validate the message
			if err := validate.Struct(MessageRequest); err != nil {
				log.Println("Validation error:", err)
				client.send <- prepareMessage(gin.H{
					"success": false,
					"message": "Error decoding message: " + err.Error(),
					"code":    422,
					"error":   err.Error(),
					"type":    "error",
				})
				continue
			}

			if MessageRequest.Type == requests_models.TypeConfig {
				if !is_host {
					client.send <- prepareMessage(gin.H{
						"success": false,
						"message": "You are not the host of the room",
						"code":    403,
						"error":   "Host only",
						"type":    "error",
					})
					break
				}
				// TODO : Change the game configuration
			}

			fmt.Printf("Message reçu de %s: %v\n", conn.RemoteAddr(), MessageRequest.Message)
			actual_room.BroadcastMessage(prepareMessage(MessageRequest), nil)

		case websocket.PongMessage:
			// Optionnel : log du pong
			log.Println("Pong reçu")

		}
	}

	// Déconnexion
	mutex.Lock()
	actual_room.RemoveClient(client) // Remove the client from the room
	cache_rooms.RemovePlayerFromRoom(user.ID, room_target.RoomID)
	delete(conn_clients, client) // Remove the client from the map
	mutex.Unlock()

	actual_room.BroadcastMessage(prepareMessage(responses_models.ResponseForSystem{
		BaseResponse: responses_models.BaseResponse{
			Code:    200,
			Type:    responses_models.TypeSystem,
			Message: fmt.Sprintf("User '%s' left the room", user.Username),
		},
		Content: gin.H{
			"status":   "left",
			"user_id":  user.ID,
			"username": user.Username,
		},
	}), map[string]bool{
		user.ID: true,
	})
	fmt.Printf("<< Client déconnecté : %s (total: %d)\n", conn.RemoteAddr(), len(conn_clients))
}
