package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	Email string             `json:"email,omitempty" bson:"email,omitempty" validate:"required"`
}
