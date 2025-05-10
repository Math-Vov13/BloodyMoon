package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin Framework!",
		})
	})

	router.POST("/", func(c *gin.Context) {
		fmt.Printf("Request Body: %s\n", c.Request.Body)
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin Framework!2",
		})
	})

	// Start the server
	router.Run(":8080")
}
