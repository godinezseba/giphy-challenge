package mappers

import (
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/infrastructure/http/giphy/dto"
)

func GIFDTOToEntity(gif *dto.GIFResponse) *entities.GIF {
	return &entities.GIF{
		ProviderID: gif.ID,
		Name:       gif.Title,
		URL:        gif.Images.Original.URL,
	}
}

func GIFDTOsToEntities(gifs []*dto.GIFResponse) []*entities.GIF {
	gifEntities := make([]*entities.GIF, len(gifs))

	for i := range gifs {
		gifEntities[i] = GIFDTOToEntity(gifs[i])
	}

	return gifEntities
}
