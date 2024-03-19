package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MakeMongoDBConnection() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURL := os.Getenv("MONGO_URL")

	log.Println("Trying to connect to MongoDB...")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))

	return client, err
}
