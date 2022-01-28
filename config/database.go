package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseCollections struct {
	Cars *mongo.Collection
}

var Collections DatabaseCollections
var Client *mongo.Client

func ConnectDatabase() error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://bundle:bundledbpassword@localhost:27011/DEV?authDb=admin"))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	db := client.Database("DEV")
	carsCollection := db.Collection("cars")
	if err != nil {
		return err
	}
	Collections = DatabaseCollections{
		Cars: carsCollection,
	}

	Client = client

	return nil
}
