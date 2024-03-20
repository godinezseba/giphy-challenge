package gif

import (
	"context"
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/domain/ports"
)

type Storage interface {
	SaveContent(ctx context.Context, gif *entities.GIF) (*entities.GIF, error)
	RemoveContent(ctx context.Context, gif *entities.GIF) (*entities.GIF, error)
	FindBySearch(ctx context.Context, search string, maxNumber int) ([]*entities.GIF, error)
}

type Persistence interface {
	SaveMetadata(ctx context.Context, gif *entities.GIF) (*entities.GIF, error)
	SaveMultipleMetadata(ctx context.Context, gifs []*entities.GIF) ([]*entities.GIF, error)
	FindByFilter(ctx context.Context, filter *entities.GIFSearch) (*entities.Pagination[entities.GIF], error)
}

type Adapter struct {
	Storage
	Persistence
}

func MakeGIFAdapter(
	PersistenceAdapter Persistence,
	StorageAdapter Storage,
) ports.GIF {
	return &Adapter{
		Storage:     StorageAdapter,
		Persistence: PersistenceAdapter,
	}
}
