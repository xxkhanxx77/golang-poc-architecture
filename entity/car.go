package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Car struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name *string            `bson:"name"`

	CreateDate *time.Time `bson:"created_at"`
	UpdateDate *time.Time `bson:"updated_at"`
}
