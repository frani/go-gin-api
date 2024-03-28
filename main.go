package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	configs "github.com/frani/go-gin-api/src/configs"
	middlewares "github.com/frani/go-gin-api/src/middlewares"
	routers "github.com/frani/go-gin-api/src/routers"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connected with database
	configs.ConnectDB()

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
