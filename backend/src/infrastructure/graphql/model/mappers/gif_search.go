package mappers

import (
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/infrastructure/graphql/model"
)

func GIFSearchModelToEntity(gifSearch *model.GIFSearchInput) *entities.GIFSearch {
	entity := &entities.GIFSearch{}

	if gifSearch.Page != nil {
		entity.Page = *gifSearch.Page
	}

	if gifSearch.PageSize != nil {
		entity.PageSize = *gifSearch.PageSize
	}

	if gifSearch.Query != nil {
		entity.Query = *gifSearch.Query
	}

	return entity
}
