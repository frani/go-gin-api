package users

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.GET("/", ListUsers)
	r.POST("/", PostUser)
	r.GET("/:id", GetUser)
	r.PATCH("/:id", PatchUser)
	r.DELETE("/:id", DeleteUser)
	return r
}
