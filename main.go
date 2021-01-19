package main

import (
	"baac-backend/handler"
	"baac-backend/repo"
	"baac-backend/services"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:MEQaC%23fJXgG@mongo:27017"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	r := gin.Default()
	uRepo := repo.NewRepo(client)
	uService := services.NewUserService(uRepo)
	handler.NewUserHandler(r, uService)

	cRepo := repo.NewCarRepo(client)
	cService := services.NewCarService(cRepo)
	handler.NewCarHandler(r, cService)
	r.Run() // listen and serve
}
