package gif

import (
	"context"
	"errors"
	gifAdapter "giphy/challenge/src/infrastructure/composite/gif"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_UseCase_GIF_Find_By_Filter_Success(t *testing.T) {
	ctx := context.Background()

	t.Run("GIFs founded", func(t *testing.T) {
		assert := assert.New(t)
		gifAdapterMock := new(gifAdapter.AdapterMock)

		gifAdapterMock.On(
			"FindByFilter",
			mock.Anything,
			mock.AnythingOfType("*entities.GIFSearch"),
		).Return(PaginationStub(), nil)

		useCase := MakeGIFUseCase(gifAdapterMock)

		foundGIFs, err := useCase.FindByFilter(ctx, GIFSearchStub())

		assert.Nil(err)
		assert.Equal(foundGIFs.Page, 1)
		assert.Equal(foundGIFs.TotalRows, 1)
		assert.Equal(foundGIFs.RowsPerPage, 1)
		assert.Equal(foundGIFs.Results[0].ID, "123")
		gifAdapterMock.AssertNumberOfCalls(t, "FindByFilter", 1)
	})
}

func Test_UseCase_GIF_Find_By_Filter_Fail(t *testing.T) {
	ctx := context.Background()

	t.Run("Error saving Content", func(t *testing.T) {
		assert := assert.New(t)
		gifAdapterMock := new(gifAdapter.AdapterMock)

		gifAdapterMock.On(
			"FindByFilter",
			mock.Anything,
			mock.AnythingOfType("*entities.GIFSearch"),
		).Return(nil, errors.New("unhandled error"))

		useCase := MakeGIFUseCase(gifAdapterMock)

		foundGIFs, err := useCase.FindByFilter(ctx, GIFSearchStub())

		assert.EqualError(err, "finding gifs error")
		assert.Nil(foundGIFs)
		gifAdapterMock.AssertNumberOfCalls(t, "FindByFilter", 1)
	})
}
