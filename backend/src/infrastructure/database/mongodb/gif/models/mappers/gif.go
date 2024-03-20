package mappers

import (
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/infrastructure/database/mongodb/gif/models"
	"time"
)

func GIFEntityToModel(gif *entities.GIF) *models.GIF {
	return &models.GIF{
		ID:        gif.ID,
		Name:      gif.Name,
		URL:       gif.URL,
		User:      gif.User,
		Tags:      gif.Tags,
		CreatedAt: time.Now().Format(time.ANSIC),
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
	gifEntities := make([]*entities.GIF, len(gifs))

	for position := range gifs {
		gifEntities[position] = GIFModelToEntity(gifs[position])
	}

	return gifEntities
}

func GIFEntitiesToInterfaces(gifs []*entities.GIF) []interface{} {
	gifModels := make([]interface{}, len(gifs))

	for position := range gifs {
		gifModels[position] = GIFEntityToModel(gifs[position])
	}

	return gifModels
}
