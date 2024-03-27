package auth

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	r.POST("/login", PostLogIn)
	r.POST("/logup", PostLogUp)
}
