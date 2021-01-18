package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        *string            `bson:"name"`
	PhoneNumber *string            `bson:"phone_number"`
	Eamil       *string            `bson:"email"`

	CreateDate *time.Time `bson:"created_at"`
	UpdateDate *time.Time `bson:"updated_at"`
}
