package configs

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB | @desc: mongo database struct
var DB *mongo.Database

// Connect | @desc: connects to mongoDB
func ConnectDB() {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		panic(err)
	}
	fmt.Println("🍀 Mongo Connected")

	// Connect to mongo database
	DB = client.Database(MONGO_DB)
}
