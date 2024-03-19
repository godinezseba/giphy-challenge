package gif

import (
	"context"
	"errors"
	"giphy/challenge/src/domain/entities"
	"log"
)

func (g GIFUseCase) Create(ctx context.Context, newGIF *entities.GIF) (*entities.GIF, error) {
	log.Println("[Use Case > GIF > Create] Storing content...")

	gifStored, err := g.GIFPersistenceAdapter.SaveContent(ctx, newGIF)
	if err != nil {
		log.Println("[Use Case > GIF > Create] Error storing content", map[string]any{"error": err.Error()})
		return nil, errors.New("creating content error")
	}

	log.Println("[Use Case > GIF > Create] Storing metadata...")

	gifStored, err = g.GIFPersistenceAdapter.SaveMetadata(ctx, gifStored)
	if err != nil {
		log.Println("[Use Case > GIF > Create] Error storing metadata", map[string]any{"error": err.Error()})

		if _, err = g.GIFPersistenceAdapter.RemoveContent(ctx, gifStored); err != nil {
			log.Println("[Use Case > GIF > Create] Error removing content", map[string]any{"error": err.Error()})
		}

		return nil, errors.New("creating metadata error")
	}

	return gifStored, nil
}
