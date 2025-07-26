package websocket_Room

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

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

const MAX_CLIENTS = 1024 * 50 // Maximum number of clients allowed in the WebSocket server

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

const (
	// writeWait = 10 * time.Second

	maxPongLoseTime     = 60 * time.Second // maximum pong wait time
	maxPongResponseTime = 5 * time.Second  // temps d'attente pour le pong

	pongWait   = 30 * time.Second
	pingPeriod = pongWait - maxPongResponseTime // ping un peu avant le timeout
)

func HandleWebSocket(c *gin.Context) {
	user := c.MustGet("user").(*users_models.User)
	excited := true

	// --- Verify Room and Player Access ---
	room_target, err1, errcode := verifyAccesstoRoom(c, user)
	if err1 != "" {
		c.Writer.Header().Set("Retry-After", "30")         // 30 seconds
		c.Writer.Header().Set("Cache-Control", "no-cache") // no cache
		c.Writer.Header().Set("Connection", "close")       // close connection
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
		log.Println("Upgrade connection error:", err)
		return
	}
	defer func() {
		if excited {
			conn.Close()
		}
	}()

	// --- Add the player to the room ---
	// Get the room
	actual_room := game_rooms[room_target.RoomID]
	if actual_room == nil {
		fmt.Println("Room not found, creating a new one")
		actual_room = CreateRoom(room_target.RoomID, room_target.HostID) // Create the room if it doesn't exist
	}

	// Check if the player is already in the room
	var is_host = (user.ID == actual_room.hostID)
	if is_host {
		cache_rooms.ChangeRoomStatus(actual_room.ID, rooms_models.StatusActive)
	} else {
		// TODO : CAUTION ! This solution does not work for AFK players
		cache_rooms.AddPlayerToRoom(user.ID, actual_room.ID)
	}
	defer func() {
		if excited {
			cache_rooms.RemovePlayerFromRoom(user.ID, actual_room.ID)
		}
	}()

	// Change the user's status to "in room"
	if succ := mongodb.ChangeUserStatus(user.ID, users_models.StatusInRoom, actual_room.ID); !succ {
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(
			websocket.CloseInternalServerErr,
			"Internal error",
		))
		return
	}

	fmt.Printf("User has joined the room: %s\n", actual_room.ID)
	fmt.Printf("User is host: %t\n", is_host)

	// Create the client
	mutex.Lock()
	client := actual_room.CreateClient(conn, user) // Create client and add it to the room
	conn_clients[client] = true                    // Add the client to the map
	mutex.Unlock()
	fmt.Printf(">> Client connected : %s (total: %d)\n", conn.RemoteAddr(), len(conn_clients))
	defer func() {
		if excited {
			actual_room.RemoveClient(client) // Remove the client from the room
			mongodb.ChangeUserStatus(user.ID, users_models.StatusOffline, "")
		}
	}()

	// --- Start the ping ticker ---
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		if excited {
			ticker.Stop() // Stop the ticker when the connection is closed
		}
	}()
	if err = conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println("Error setting read deadline:", err)
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(
			websocket.CloseInternalServerErr,
			"Internal error",
		))
		return
	}
	// Set the ping handler to update the last ping time
	conn.SetPongHandler(func(appData string) error {
		log.Println("pong received")
		mutex.Lock()
		client.lastPing = time.Now().Unix() // Update the last ping time
		mutex.Unlock()
		//ticker.Reset(pingPeriod) // Reset the ticker
		if err = conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			log.Println("Error setting read deadline:", err)
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(
				websocket.CloseInternalServerErr,
				"Internal error",
			))
			return fmt.Errorf("error setting read deadline: %w", err)
		}
		return nil
	})

	go func() {
		for msg := range client.send {
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Printf("Writting Message error: '%v'\n", err)
				break
			}
		}
		fmt.Println("Client send channel closed, stopping goroutine")
	}()
	go func() {
		for range ticker.C {
			// Envoie le ping
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("Error sending Ping: '%v'\n", err)
				break
			}

			log.Println("Ping sent!")
		}
		fmt.Println("Ticker stopped!")
	}()

	// --- Read messages from the client ---
	go func() {
		defer func(actual_room *Room) {
			defer func() {
				// Choose a new host if the current host is leaving
				if len(actual_room.clients) == 0 {
					fmt.Printf("No more clients: deleting Room (%s)\n", actual_room.ID)
					fmt.Printf("Deleted: %v\n\n", cache_rooms.DeleteRoom(actual_room.ID)) // Delete the room if no clients are left
					mutex.Lock()
					actual_room.RemoveRoom() // Remove the room from the map
					mutex.Unlock()
				} else {
					if user.ID == actual_room.hostID {
						// TODO: Change the host if the current host is leaving
						newHost, err := actual_room.ChangeRoomHost()
						if err != nil {
							fmt.Printf("Error changing host: %v\n", err)
						} else {
							fmt.Printf("New host for room %s is %s\n", actual_room.ID, newHost.ID)
							// Notify all clients about the new host
							cache_rooms.ChangeRoomHost(actual_room.ID, newHost.ID)
							actual_room.BroadcastMessage(prepareMessage(responses_models.ResponseForSystem{
								BaseResponse: responses_models.BaseResponse{
									Code:    200,
									Type:    responses_models.TypeSystem,
									Message: fmt.Sprintf("New Host '%s' choosed for the room", newHost.Username),
								},
								Content: gin.H{
									"status":   "host_changed",
									"user_id":  newHost.ID,
									"username": newHost.Username,
								},
							}), nil)
						}
					}
				}
			}()
			ticker.Stop()
			// Déconnexion
			mutex.Lock()
			actual_room.RemoveClient(client) // Remove the client from the room
			mutex.Unlock()
			cache_rooms.RemovePlayerFromRoom(user.ID, actual_room.ID)
			mongodb.ChangeUserStatus(user.ID, users_models.StatusOffline, "")

			mutex.Lock()
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
			fmt.Printf("<< Client disconnected : %s (total: %d)\n", conn.RemoteAddr(), len(conn_clients))
		}(actual_room)

		// Read Messages from the client
		for {
			mt, msg, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Reading Message error: '%v'", err)
				conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(
					websocket.CloseMessage,
					"Ping Timeout or Read Error",
				))
				return
			}
			// conn.SetWriteDeadline(time.Now().Add(writeWait))
			conn.SetReadLimit(1024 * 5) // Set a limit for the message size

			switch mt {
			case websocket.TextMessage:
				reqMessage, code, err := decodeMessage(msg)
				if err != nil {
					log.Println("Message decode error:", err)
					client.send <- prepareMessage(gin.H{
						"success": false,
						"message": "Error in Message validation",
						"code":    code,
						"error":   err.Error(),
						"type":    "error",
					})
					continue
				}

				if reqMessage.Action == requests_models.TypeConfig {
					if !(user.ID == actual_room.hostID) {
						client.send <- prepareMessage(gin.H{
							"success": false,
							"message": "You are not the host of the room",
							"code":    403,
							"error":   "Host only",
							"type":    "error",
						})
						continue
					}
					// TODO : Change the game configuration
				}

				fmt.Printf("Message received from %s: '%v'\n", conn.RemoteAddr(), reqMessage.Message)
				actual_room.BroadcastMessage(prepareMessage(reqMessage), nil)

			default:
				log.Printf("Received non-text message type: %d\n", mt)
				conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(
					websocket.CloseUnsupportedData,
					"Only text messages are supported",
				))
				return
			}
		}
	}()

	// Say Welcome to the client !
	// --> client
	client.send <- prepareMessage(gin.H{
		"success": true,
		"message": "You are now connected to the room!",
		"room":    actual_room,
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

	excited = false
}
