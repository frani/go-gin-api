package users

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.GET("/", ListUsers)
	r.POST("/", PostUser)
	r.GET("/:id", GetUser)
	// r.PATCH("/:userId", PatchUser)
	// r.DELETE("/:userId", DeleteUser)
	return r
}
