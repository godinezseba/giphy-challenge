package adapters

import (
	"context"
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/infrastructure/database/mongodb/gif/models/mappers"
	"log"
)

func (g *gifMongoDBAdapter) SaveMultipleMetadata(ctx context.Context, gifs []*entities.GIF) ([]*entities.GIF, error) {
	log.Println("[MongoDB > GIF > SaveMultipleMetadata] Saving...")

	gifModels := mappers.GIFEntitiesToInterfaces(gifs)

	_, err := g.gifCollection.InsertMany(ctx, gifModels)

	if err != nil {
		log.Println("[MongoDB > GIF > SaveMultipleMetadata] Error saving")

		return nil, err
	}

	return gifs, nil
}
