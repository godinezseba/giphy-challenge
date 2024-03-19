package adapters

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type GIFMongoDBAdapter struct {
	gifCollection *mongo.Collection
}

func MakeBookMongoDBAdapter(client *mongo.Client) *GIFMongoDBAdapter {
	mongoGIFDatabase := os.Getenv("MONGO_GIF_DATABASE")
	mongoGIFCollection := os.Getenv("MONGO_GIF_COLLECTION")

	db := client.Database(mongoGIFDatabase)
	gifCollection := db.Collection(mongoGIFCollection)

	return &GIFMongoDBAdapter{
		gifCollection,
	}
}
