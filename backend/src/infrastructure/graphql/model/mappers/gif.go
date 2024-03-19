package mappers

import (
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/infrastructure/graphql/model"
)

func GIFInputModelToEntity(gif *model.GIFInput) *entities.GIF {
	return &entities.GIF{
		Name:    gif.Name,
		User:    gif.User,
		Tags:    gif.Tags,
		Content: gif.Content,
	}
}

func GIFEntityToModel(gif *entities.GIF) *model.Gif {
	return &model.Gif{
		ID:        gif.ID,
		Name:      gif.Name,
		URL:       gif.URL,
		User:      gif.User,
		Tags:      gif.Tags,
		CreatedAt: gif.CreatedAt.GoString(),
	}
}

func GIFEntitiesToModels(gifs []*entities.GIF) []*model.Gif {
	gifModels := make([]*model.Gif, len(gifs))

	for position := range gifs {
		gifModels[position] = GIFEntityToModel(gifs[position])
	}

	return gifModels
}

func GifPaginationEntityToModel(
	pagination *entities.Pagination[entities.GIF],
) *model.GIFPagination {
	return &model.GIFPagination{
		Page:        pagination.Page,
		RowsPerPage: pagination.RowsPerPage,
		TotalRows:   pagination.TotalRows,
		Results:     GIFEntitiesToModels(pagination.Results),
	}
}
