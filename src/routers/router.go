package routers

import (
	AuthRouter "github.com/frani/go-gin-api/src/routers/auth"
	StatusRouter "github.com/frani/go-gin-api/src/routers/status"
	UsersRouter "github.com/frani/go-gin-api/src/routers/users"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	AuthRouter.InitRouter(r.Group("/auth"))
	StatusRouter.InitRouter(r.Group("/status"))
	UsersRouter.InitRouter(r.Group("/users"))
}
