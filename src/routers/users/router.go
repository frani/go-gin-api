package users

import "github.com/gin-gonic/gin"

func UsersRoutes() {
	r := gin.Default()
	route := r.Group("/")
	route.GET("/", ListUsers)
	route.POST("/", PostUser)
	route.GET("/:userId", GetUser)
	route.PATCH("/:userId", PatchUser)
	route.DELETE("/:userId", DeleteUser)
}
