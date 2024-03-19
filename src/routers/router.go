package routers

import (
	StatusRouter "github.com/frani/go-gin-api/src/routers/status"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	StatusRouter.InitRouter(r.Group("/status"))
}
