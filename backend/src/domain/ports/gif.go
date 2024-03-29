package ports

import (
	"context"
	"giphy/challenge/src/domain/entities"
)

type GIF interface {
	SaveContent(ctx context.Context, gif *entities.GIF) (*entities.GIF, error)
	SaveMetadata(ctx context.Context, gif *entities.GIF) (*entities.GIF, error)
	SaveMultipleMetadata(ctx context.Context, gifs []*entities.GIF) ([]*entities.GIF, error)
	FindByFilter(ctx context.Context, filter *entities.GIFSearch) (*entities.Pagination[entities.GIF], error)
	RemoveContent(ctx context.Context, gif *entities.GIF) (*entities.GIF, error)
	FindBySearch(ctx context.Context, search string, maxNumber int) ([]*entities.GIF, error)
}
