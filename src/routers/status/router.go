package status

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.GET("/", GetStatus)
	return r
}
