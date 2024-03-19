package mappers

import (
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/infrastructure/graphql/model"
)

func GIFSearchModelToEntity(gifSearch *model.GIFSearchInput) *entities.GIFSearch {
	entity := &entities.GIFSearch{
		Page: gifSearch.Page,
	}

	if gifSearch.Query != nil {
		entity.Query = *gifSearch.Query
	}

	return entity
}
