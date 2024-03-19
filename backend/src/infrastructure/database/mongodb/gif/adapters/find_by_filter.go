package adapters

import (
	"context"
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/infrastructure/database/mongodb/gif/models"
	"giphy/challenge/src/infrastructure/database/mongodb/gif/models/mappers"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (g *gifMongoDBAdapter) FindByFilter(ctx context.Context, filter *entities.GIFSearch) (*entities.Pagination[entities.GIF], error) {
	log.Println("[MongoDB > GIF > FindByFilter] Searching...")

	var gifModels []*models.GIF

	cursor, err := g.gifCollection.Find(ctx, bson.D{
		primitive.E{
			Key:   "name",
			Value: filter.Query,
		},
	})

	if err != nil {
		log.Println("[MongoDB > GIF > FindByFilter] Error finding results")

		return nil, err
	}

	if err := cursor.All(ctx, gifModels); err != nil {
		log.Println("[MongoDB > GIF > FindByFilter] Error converting results")

		return nil, err
	}

	return &entities.Pagination[entities.GIF]{
		Page:    filter.Page,
		Results: mappers.GIFModelsToEntities(gifModels),
	}, nil
}
