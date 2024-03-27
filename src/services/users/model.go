package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty" validate:"required,min=3,max=60"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Password string             `json:"password,omitempty" bson:"password,omitempty" validate:"required,password,min=10,max=60"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty" validate:"required,min=3,max=60"`
	Lastname string             `json:"lastname,omitempty" bson:"lastname,omitempty" validate:"required,min=3,max=60"`
	Roles    []string           `json:"roles" bson:"roles" default:"['user']" validate:"required,min=3,max=60"`
}
