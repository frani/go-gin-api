package status

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get Status OK
// @Success 200
// @Router /status [get]

func GetStatus(ctx *gin.Context) {
	// Return status 200 OK.
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    nil,
		"message": "ok",
		"errors":  nil,
	})
}
