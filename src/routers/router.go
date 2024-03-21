package routers

import (
	StatusRouter "github.com/frani/go-gin-api/src/routers/status"
	UsersRouter "github.com/frani/go-gin-api/src/routers/users"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	StatusRouter.InitRouter(r.Group("/status"))
	UsersRouter.InitRouter(r.Group("/users"))
}
