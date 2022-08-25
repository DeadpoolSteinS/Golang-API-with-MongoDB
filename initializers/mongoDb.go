package initializers

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectToDB() {
	// get database url from .env file
	dns := os.Getenv("DB_URL")

	clientOptions := options.Client().ApplyURI(dns)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	// get/create products collections
	// Users = client.Database("practice_1").Collection("users")
}
