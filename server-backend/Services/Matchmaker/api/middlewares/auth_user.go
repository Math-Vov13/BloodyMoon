package middlewares

import (
	"net/http"
	"os"

	"github.com/Math-Vov13/BloodyMoon/internal/database/mongodb"
	"github.com/Math-Vov13/BloodyMoon/models/users_models"
	"github.com/gin-gonic/gin"
)

func VerifyUserMiddleware() gin.HandlerFunc {
	if os.Getenv("environment") != "production" {
		return func(ctx *gin.Context) {
			username := ctx.Query("testName")
			if username == "" {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"env":     "development",
					"message": "You must provide a username!",
				})
				ctx.Abort()
				return
			}

			user_db := mongodb.GetUserByName(username)
			if user_db == nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"env":     "development",
					"message": "You must provide a valid username!",
				})
				ctx.Abort()
				return
			}

			ctx.Set("user", &users_models.User{
				ID:       user_db.ID,
				Username: user_db.Username,
			})
			ctx.Next()
		}
	}

	return func(ctx *gin.Context) {
		// Check if the user is authenticated
		if ctx.GetHeader("Authorization") == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
	}
}
