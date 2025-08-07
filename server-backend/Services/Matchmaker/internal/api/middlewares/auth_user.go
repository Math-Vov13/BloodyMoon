package middlewares

import (
	"net/http"
	"os"

	"github.com/Math-Vov13/BloodyMoon/internal/database/cache_redis/cache_sessions"
	"github.com/gin-gonic/gin"
)

func VerifyUserMiddleware() gin.HandlerFunc {
	if os.Getenv("environment") != "production" {
		return func(ctx *gin.Context) {
			username := ctx.Query("testName")
			if username == "" {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"env":     "development",
					"message": "You must provide a username! (key: testName)",
				})
				ctx.Abort()
				return
			}

			user_db, err := cache_sessions.GetUserByName(username)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"env":     "development",
					"message": "You must provide a valid username! (key: testName)",
				})
				ctx.Abort()
				return
			}

			ctx.Set("user", user_db)
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
