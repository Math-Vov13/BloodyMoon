package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Math-Vov13/BloodyMoon/api/endpoints"
	"github.com/Math-Vov13/BloodyMoon/api/middlewares"
	"github.com/Math-Vov13/BloodyMoon/api/websocket_Queue"
	"github.com/Math-Vov13/BloodyMoon/api/websocket_Room"
	"github.com/Math-Vov13/BloodyMoon/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	router := gin.Default()
	config.ConfigRuntime()

	//------ Rest API ------//
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "[MatchMaker Service]: Welcome to BloodyMoon!",
		})
	})

	games_endpoint := router.Group("/games", middlewares.VerifyUserMiddleware())
	games_endpoint.POST("/create", endpoints.CreatePrivateGame)
	games_endpoint.POST("/delete", endpoints.DeletePrivateGame)

	// STATIC FILES
	router.StaticFile("/home", "../../web/templates/index.html")
	router.Static("/static", "../../web/static")

	//------ WEBSOCKETS ------//

	// Queue WebSocket
	// RateLimit : 3/10s
	router.GET("/ws/queue", middlewares.VerifyUserMiddleware(), websocket_Queue.HandleWebSocket)

	// Game WebSocket
	// RateLimit : 3/10s
	router.GET("/ws/room", middlewares.VerifyUserMiddleware(), websocket_Room.HandleWebSocket)

	//------ RUN SERVER ------//
	// Start the server
	port := "8001"
	if envPort := os.Getenv("Port"); envPort != "" {
		port = envPort
	}
	fmt.Printf("\n> Server is running on %s:%s\n", os.Getenv("HOST"), port)
	router.Run(os.Getenv("HOST") + ":" + port)
}
