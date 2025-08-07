package middlewares

import (
	"net/http"

	"github.com/Math-Vov13/BloodyMoon/internal/database/cache_redis/cache_rooms"
	"github.com/Math-Vov13/BloodyMoon/models/cache/rooms_models"
	"github.com/gin-gonic/gin"
)

func VerifyRoomCodeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roomCode := ctx.Query("code")
		if roomCode == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Room code is required",
			})
			ctx.Abort()
			return
		}

		if len(roomCode) != 6 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid room code format",
			})
			ctx.Abort()
			return
		}

		room, err := cache_rooms.GetRoomByCode(roomCode)
		if err != nil || room == nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "Room not found",
			})
			ctx.Abort()
			return
		}

		if room.Status != rooms_models.StatusActive {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "Room is not active",
			})
			ctx.Abort()
			return
		}

		ctx.Set("room", room)
		ctx.Next()
	}
}
