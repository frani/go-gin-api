package users

import (
	"context"

	configs "github.com/frani/go-gin-api/src/configs"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserCollection | @desc: the user ccollection on the database
var UserCollection *mongo.Collection

/*
CreateUserSchema
@desc: adds schema validation and indexes to collection
*/
func CreateUser(Name string, Lastname string, Password string, Email string, Username string) (result interface{}, err error) {

	newUser := User{
		Id:       primitive.NewObjectID(),
		Name:     Name,
		Lastname: Lastname,
		Password: Password,
		Username: Username,
		Email:    Email,
	}

	result, err = configs.DB.Collection("users").InsertOne(context.TODO(), newUser)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// func ListUsers(filter bson) error {
// 	cursor, err := UserCollection.Find(configs.Ctx, filter)

// 	if err != nil {
// 		return err
// 	}

// 	defer cursor.Close(configs.Ctx)

// 	for cursor.Next(configs.Ctx) {
// 		var user User

// 		err := cursor.Decode(&user)

// 		if err != nil {
// 			return err
// 		}

// 		fmt.Println(user)
// 	}

// 	return users
// }
