package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RateLimit(maxRequests int, duration time.Duration) gin.HandlerFunc {
	requests := make(map[string]int)

	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()

		if count, exists := requests[ip]; exists {

			if count >= maxRequests {
				ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
					"error":   "too many requests",
					"success": false,
					"message": "too many requests",
				})
				return
			}

			requests[ip]++
		} else {

			requests[ip] = 1
		}

		time.AfterFunc(duration, func() {
			delete(requests, ip)
		})

		ctx.Next()
	}
}
