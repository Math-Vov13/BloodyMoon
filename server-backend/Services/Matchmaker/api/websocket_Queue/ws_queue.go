package websocket_Queue

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // autorise toutes les origines
	},
}

var serviceAvailable = false
var queue = make(chan string, 100)           // Channel for queue messages
var clients = make(map[*websocket.Conn]bool) // Connected clients

type status struct {
	Connected       bool `json:"connected"`
	HeartBeat       int  `json:"heartbeat"`
	NumberOfClients int  `json:"number_of_clients"`
}

func HandleWebSocket(c *gin.Context) {
	if !serviceAvailable {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"success": false,
			"message": "Sorry, Queue Service is not available for the moment!",
			"code":    http.StatusServiceUnavailable,
			"type":    "error",
		})
		return
	}

	// conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	// if err != nil {
	// 	log.Println("Erreur WebSocket :", err)
	// 	return
	// }

	// // NOT IMPLEMENTED
	// conn.WriteJSON(gin.H{
	// 	"success": false,
	// 	"message": "Sorry, Queue Service is not accessible for the moment!",
	// 	"type":    "error",
	// 	"code":    500,
	// })
	// conn.Close()
}

func GetSocketStatus() status {
	return status{
		Connected:       false,
		NumberOfClients: len(clients),
	}
}
