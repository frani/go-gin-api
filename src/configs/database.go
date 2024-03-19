package configs

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Ctx | @desc: mongo context interface
var Ctx context.Context

// Cancel | @desc: mongo context cancel function
var Cancel context.CancelFunc

// Client | @desc: mongo client struct
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

	Ctx, Cancel = context.WithTimeout(context.Background(), 30*time.Second)
	client, err = mongo.Connect(Ctx, options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		panic(err)
	}

	// Connect to mongo database
	DB = client.Database(MONGO_DB)
}
