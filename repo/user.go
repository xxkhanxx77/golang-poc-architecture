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

type Repo struct {
	client *mongo.Client
}

func NewRepo(c *mongo.Client) *Repo {
	return &Repo{client: c}
}

func (r *Repo) User() (*entity.User, error) {

	u := entity.User{
		ID:          primitive.NewObjectID(),
		Name:        toPStr("xx"),
		PhoneNumber: toPStr("095"),
		Email:       toPStr("kk@gmail.com")}
	collection := r.client.Database("baacbanking").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, u)
	if err != nil {
		return nil, err
	}

	var result entity.User
	filter := bson.M{"name": "xx"}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return &result, nil
}

func toPStr(v string) *string {
	return &v
}

func (r *Repo) CreateUser(user *entity.User) (*entity.User, error) {
	collection := r.client.Database("baacbanking").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
