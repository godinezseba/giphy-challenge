package adapters

import (
	"context"
	"fmt"
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/infrastructure/database/mongodb/gif/models"
	"giphy/challenge/src/infrastructure/database/mongodb/gif/models/mappers"
	"log"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
)

func (g *gifMongoDBAdapter) FindByFilter(ctx context.Context, filter *entities.GIFSearch) (*entities.Pagination[entities.GIF], error) {
	log.Println("[MongoDB > GIF > FindByFilter] Searching...")

	searchString := filter.Query
	query := bson.M{}

	if searchString != "" {
		pattern := fmt.Sprintf("(?i)%s", regexp.QuoteMeta(searchString))

		query = bson.M{
			"$or": []bson.M{
				{"name": bson.M{"$regex": pattern}},
				{"tags": bson.M{"$in": []string{searchString}}},
			},
		}
	}

	cursor, err := g.gifCollection.Find(ctx, query)

	if err != nil {
		log.Println("[MongoDB > GIF > FindByFilter] Error finding results")

		return nil, err
	}

	gifModels := make([]*models.GIF, cursor.RemainingBatchLength())

	if err := cursor.All(ctx, &gifModels); err != nil {
		log.Println("[MongoDB > GIF > FindByFilter] Error converting results")

		return nil, err
	}

	return &entities.Pagination[entities.GIF]{
		Page:    filter.Page,
		Results: mappers.GIFModelsToEntities(gifModels),
	}, nil
}
