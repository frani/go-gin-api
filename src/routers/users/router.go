package users

import (
	"github.com/frani/go-gin-api/src/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	r.Use(middlewares.Authorize("admin"))
	r.GET("/", ListUsers)
	r.POST("/", PostUser)
	r.GET("/:id", GetUser)
	r.PATCH("/:id", PatchUser)
	r.DELETE("/:id", DeleteUser)
}
