package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var MONGO_DB string
var MONGO_URI string
var JWT_SECRET string
var JWT_EXPIRY_TIME string
var JWT_ISSUER string

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// MONGO
	MONGO_URI = os.Getenv("MONGO_URI")
	MONGO_DB = os.Getenv("MONGO_DB")

	// JWT
	JWT_SECRET = os.Getenv("JWT_SECRET")
	JWT_EXPIRY_TIME = os.Getenv("JWT_EXPIRY_TIME")
	JWT_ISSUER = os.Getenv("JWT_ISSUER")

}
