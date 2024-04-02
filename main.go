package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	docs "github.com/frani/go-gin-api/docs"
	configs "github.com/frani/go-gin-api/src/configs"
	middlewares "github.com/frani/go-gin-api/src/middlewares"
	routers "github.com/frani/go-gin-api/src/routers"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Create Gin App
	r := gin.New()

	r.SetTrustedProxies(nil)

	// Middleware
	r.Use(gin.Logger())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(middlewares.RateLimit(10, time.Second))
	r.Use(gin.Recovery())

	// Bind routes
	routers.InitRouters(r)

	// Docs
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/docs", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/docs/index.html")
	})

	// Handle not founds
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	// Listen on port 8080
	fmt.Println("🌥️  Listen on port ", configs.PORT)
	err := r.Run(configs.PORT)
	if err != nil {
		log.Fatal(err)

	}
}
