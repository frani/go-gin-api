package configs

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// client | @desc: mongo client struct
var client *mongo.Client

// DB | @desc: mongo database struct
var DB *mongo.Database

// Connect | @desc: connects to mongoDB
func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MONGO_URI := os.Getenv("MONGO_URI")
	MONGO_DB := os.Getenv("MONGO_DB")

	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		panic(err)
	}
	fmt.Println("🍀 Mongo Connected")

	// Connect to mongo database
	DB = client.Database(MONGO_DB)
}
