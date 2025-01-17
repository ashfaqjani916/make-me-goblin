
package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() {




	// Correct way to create ClientOptions using pointer
  MONGODB_URL := "mongodb+srv://ashfaqjani916:hellohs72oi@cluster0.wi6u3.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	clientOptions := options.Client().ApplyURI(MONGODB_URL)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	log.Println("Connected to MongoDB successfully!")
}
