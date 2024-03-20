package gif

import (
	"context"
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/domain/ports"
)

type GIFUseCaseInterface interface {
	Create(ctx context.Context, newGIF *entities.GIF) (*entities.GIF, error)
	FindByFilter(
		ctx context.Context,
		query *entities.GIFSearch,
	) (*entities.Pagination[entities.GIF], error)
	Populate(ctx context.Context)
}

type GIFUseCase struct {
	GIFPersistenceAdapter ports.GIF
}

func MakeGIFUseCase(
	gifPersistenceAdapter ports.GIF,
) GIFUseCaseInterface {
	return &GIFUseCase{
		GIFPersistenceAdapter: gifPersistenceAdapter,
	}
}
