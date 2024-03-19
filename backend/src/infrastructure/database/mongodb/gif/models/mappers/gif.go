package mappers

import (
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/infrastructure/database/mongodb/gif/models"
)

func GIFEntityToModel(gif *entities.GIF) *models.GIF {
	return &models.GIF{
		ID:        gif.ID,
		Name:      gif.Name,
		URL:       gif.URL,
		User:      gif.User,
		Tags:      gif.Tags,
		CreatedAt: gif.CreatedAt.GoString(),
	}
}

func GIFModelToEntity(gif *models.GIF) *entities.GIF {
	return &entities.GIF{
		ID:   gif.ID,
		Name: gif.Name,
		URL:  gif.URL,
		User: gif.User,
		Tags: gif.Tags,
	}
}

func GIFModelsToEntities(gifs []*models.GIF) []*entities.GIF {
	gifModels := make([]*entities.GIF, len(gifs))

	for position := range gifs {
		gifModels[position] = GIFModelToEntity(gifs[position])
	}

	return gifModels
}
