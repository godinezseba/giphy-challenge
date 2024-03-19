package adapters

import (
	"context"
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/infrastructure/database/mongodb/gif/models/mappers"
	"log"
)

func (g *gifMongoDBAdapter) SaveMetadata(ctx context.Context, gif *entities.GIF) (*entities.GIF, error) {
	log.Println("[MongoDB > GIF > SaveMetadata] Saving...")

	gifModel := mappers.GIFEntityToModel(gif)

	insertID, err := g.gifCollection.InsertOne(ctx, gifModel)

	if err != nil {
		log.Println("[MongoDB > GIF > SaveMetadata] Error saving")

		return nil, err
	}

	gif.ID = insertID.InsertedID.(string)

	return gif, nil
}
