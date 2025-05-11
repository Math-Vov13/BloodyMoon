package main

import (
	"fmt"
	"net/http"

	"github.com/Math-Vov13/BloodyMoon/internal/config"
	fakedb "github.com/Math-Vov13/BloodyMoon/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "[Auth Service]: Welcome to BloodyMoon!",
		})
	})

	router.POST("/register", func(c *gin.Context) {
		var body config.UserCreate
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if body.Username == "" || body.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Username and password are required",
			})
			return
		}

		fmt.Printf("Formulaire : %v\n", body)

		// Simulate user registration logic
		if !fakedb.AddUser(body.Username, body.Password) {
			c.JSON(http.StatusConflict, gin.H{
				"error": "User already exists",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User registered successfully!",
		})
	})

	router.POST("/login", func(c *gin.Context) {
		// Simulate user login logic
		var body config.UserCreate
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if body.Username == "" || body.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Username and password are required",
			})
			return
		}

		fmt.Printf("Formulaire : %v\n", body)

		if fakedb.GetUser(body.Username, body.Password) {
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Welcome back, %s!", body.Username),
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
		}
	})

	// Start the server
	router.Run(":8080")
}
