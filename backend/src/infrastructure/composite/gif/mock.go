package gif

import (
	"context"
	"giphy/challenge/src/domain/entities"

	"github.com/stretchr/testify/mock"
)

type AdapterMock struct {
	mock.Mock
}

func (a *AdapterMock) SaveContent(ctx context.Context, gif *entities.GIF) (newGIF *entities.GIF, err error) {
	args := a.Called(ctx, gif)

	if args.Get(0) != nil {
		newGIF = args.Get(0).(*entities.GIF)
	}

	if args.Get(1) != nil {
		err = args.Error(1)
	}

	return newGIF, err
}

func (a *AdapterMock) SaveMetadata(ctx context.Context, gif *entities.GIF) (newGIF *entities.GIF, err error) {
	args := a.Called(ctx, gif)

	if args.Get(0) != nil {
		newGIF = args.Get(0).(*entities.GIF)
	}

	if args.Get(1) != nil {
		err = args.Error(1)
	}

	return newGIF, err
}

func (a *AdapterMock) SaveMultipleMetadata(ctx context.Context, gif []*entities.GIF) (newGIF []*entities.GIF, err error) {
	args := a.Called(ctx, gif)

	if args.Get(0) != nil {
		newGIF = args.Get(0).([]*entities.GIF)
	}

	if args.Get(1) != nil {
		err = args.Error(1)
	}

	return newGIF, err
}

func (a *AdapterMock) FindByFilter(
	ctx context.Context,
	filter *entities.GIFSearch,
) (pagination *entities.Pagination[entities.GIF], err error) {
	args := a.Called(ctx, filter)

	if args.Get(0) != nil {
		pagination = args.Get(0).(*entities.Pagination[entities.GIF])
	}

	if args.Get(1) != nil {
		err = args.Error(1)
	}

	return pagination, err
}

func (a *AdapterMock) FindBySearch(
	ctx context.Context, search string, maxNumber int,
) (gifs []*entities.GIF, err error) {
	args := a.Called(ctx, search)

	if args.Get(0) != nil {
		gifs = args.Get(0).([]*entities.GIF)
	}

	if args.Get(1) != nil {
		err = args.Error(1)
	}

	return gifs, err
}

func (a *AdapterMock) RemoveContent(ctx context.Context, gif *entities.GIF) (newGIF *entities.GIF, err error) {
	args := a.Called(ctx, gif)

	if args.Get(0) != nil {
		newGIF = args.Get(0).(*entities.GIF)
	}

	if args.Get(1) != nil {
		err = args.Error(1)
	}

	return newGIF, err
}
