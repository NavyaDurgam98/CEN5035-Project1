package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Email     string             `bson:"email"`
	DOB       string             `bson:"dob"`
	Gender    string             `bson:"gender"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
}
