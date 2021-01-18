package repo

import (
	"baac-backend/entity"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CarRepo struct {
	client *mongo.Client
}

func NewCarRepo(c *mongo.Client) *Repo {
	return &Repo{client: c}
}

func (r *Repo) Car() (*entity.Car, error) {

	u := entity.Car{
		ID:   primitive.NewObjectID(),
		Name: toPStr("audi"),
	}
	collection := r.client.Database("baacbanking").Collection("cars")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, u)
	if err != nil {
		return nil, err
	}

	var result entity.Car
	filter := bson.M{"name": "audi"}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return &result, nil
}
