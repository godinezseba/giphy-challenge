package adapters

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type gifMongoDBAdapter struct {
	gifCollection *mongo.Collection
}

func MakeGIFMongoDBAdapter(client *mongo.Client) *gifMongoDBAdapter {
	mongoGIFDatabase := os.Getenv("MONGO_GIF_DATABASE")
	mongoGIFCollection := os.Getenv("MONGO_GIF_COLLECTION")

	db := client.Database(mongoGIFDatabase)
	gifCollection := db.Collection(mongoGIFCollection)

	return &gifMongoDBAdapter{
		gifCollection,
	}
}
