package util

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
	dbName                   = "tracker"
	bookCol                  = "book"
	noteCol                  = "note"
)

type Database struct {
	Client *mongo.Client
	Handle *mongo.Database
}

func NewDatabase(mongoURI string) (Database, context.Context) {

	dbClient, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	err = dbClient.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to cluster: %v\n", err)
	} else {
		fmt.Println("Connected to MongoDB...")
	}

	db := Database{
		Client: dbClient,
		Handle: dbClient.Database(dbName),
	}
	return db, ctx
}

func getConnection(mongoURI string) (*mongo.Client, context.Context, context.CancelFunc) {

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}
