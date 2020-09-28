package gweb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

// GetConnection - Retrieves a client to the DocumentDB
func getConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	// username := os.Getenv("MONGODB_USERNAME")
	// password := os.Getenv("MONGODB_PASSWORD")
	// clusterEndpoint := os.Getenv("MONGODB_ENDPOINT")
	// username := "root"
	// password := "toor"
	// clusterEndpoint := "localhost:27010"

	// connectionURI := fmt.Sprintf(connectionStringTemplate, username, password, clusterEndpoint)

	devURI := "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&ssl=false"
	// client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	client, err := mongo.NewClient(options.Client().ApplyURI(devURI))
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
