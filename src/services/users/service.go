package users

import (
	"context"

	configs "github.com/frani/go-gin-api/src/configs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func List(filter bson.M) (result bson.M, err error) {

	// Get all users.
	cursor, err := configs.DB.Collection("users").Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var users []bson.M
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}

	return result, nil
}

/*
CreateUserSchema
@desc: adds schema validation and indexes to collection
*/
func CreateOne(Name string, Lastname string, Password string, Email string, Username string) (result interface{}, err error) {

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

func FindOne(filter bson.M) (result bson.M, err error) {

	err = configs.DB.Collection("users").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateOne(filter bson.M, update bson.M) (result bson.M, err error) {

	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = configs.DB.Collection("users").FindOneAndUpdate(context.TODO(), filter, update, opt).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteOne(filter bson.M) (result bson.M, err error) {

	err = configs.DB.Collection("users").FindOneAndDelete(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
