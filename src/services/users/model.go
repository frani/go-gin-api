package users

import (
	"context"

	configs "github.com/frani/go-gin-api/src/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty" validate:"required,min=3,max=60"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Password string             `json:"password,omitempty" bson:"password,omitempty" validate:"required,password,min=10,max=60"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=3,max=60"`
	Lastname string             `json:"lastname,omitempty" bson:"lastname,omitempty" validate:"required,min=3,max=60"`
	Roles    []string           `json:"roles" bson:"roles" default:"['user']" validate:"required,min=3,max=60"`
}

func init() {

	index := mongo.IndexModel{
		Keys: bson.D{
			{Key: "email", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	}

	configs.DB.Collection("users").Indexes().CreateOne(context.TODO(), index)
}
