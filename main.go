package main

import (
	"log"
	"net/http"

	configs "github.com/frani/go-gin-api/src/configs"
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
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// Bind routes
	routers.InitRouters(r)

	// Handle not founds
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	// Listen on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
