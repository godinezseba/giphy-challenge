package gif

import (
	"context"
	"errors"
	gifAdapter "giphy/challenge/src/infrastructure/composite/gif"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_UseCase_GIF_Create_Success(t *testing.T) {
	ctx := context.Background()

	t.Run("GIF created", func(t *testing.T) {
		assert := assert.New(t)
		gifAdapterMock := new(gifAdapter.AdapterMock)

		gifAdapterMock.On(
			"SaveContent",
			mock.Anything,
			mock.AnythingOfType("*entities.GIF"),
		).Return(GIFWithURLStub(), nil)

		gifAdapterMock.On(
			"SaveMetadata",
			mock.Anything,
			mock.AnythingOfType("*entities.GIF"),
		).Return(GIFStub(), nil)

		useCase := MakeGIFUseCase(gifAdapterMock)

		createdGIF, err := useCase.Create(ctx, GIFToBeCreatedStub())

		assert.Nil(err)
		assert.Equal(createdGIF.ID, "123")
		assert.Equal(createdGIF.Name, "name")
		assert.Equal(createdGIF.URL, "http//google.com")
		assert.Equal(createdGIF.Content, "content")
		assert.Equal(createdGIF.User, "user")
		assert.Equal(createdGIF.Tags, []string{"tag1", "tag2"})
		assert.Equal(createdGIF.CreatedAt, &time.Time{})
		gifAdapterMock.AssertNumberOfCalls(t, "SaveContent", 1)
		gifAdapterMock.AssertNumberOfCalls(t, "SaveMetadata", 1)
		gifAdapterMock.AssertNumberOfCalls(t, "RemoveContent", 0)
	})
}

func Test_UseCase_GIF_Create_Fail(t *testing.T) {
	ctx := context.Background()

	t.Run("Error saving Content", func(t *testing.T) {
		assert := assert.New(t)

		gifAdapterMock := new(gifAdapter.AdapterMock)

		gifAdapterMock.On(
			"SaveContent",
			mock.Anything,
			mock.AnythingOfType("*entities.GIF"),
		).Return(nil, errors.New("unhandled error"))

		useCase := MakeGIFUseCase(gifAdapterMock)

		createdGIF, err := useCase.Create(ctx, GIFToBeCreatedStub())

		assert.Nil(createdGIF)
		assert.EqualError(err, "creating content error")
		gifAdapterMock.AssertNumberOfCalls(t, "SaveContent", 1)
		gifAdapterMock.AssertNumberOfCalls(t, "SaveMetadata", 0)
		gifAdapterMock.AssertNumberOfCalls(t, "RemoveContent", 0)
	})

	t.Run("Error saving Metadata", func(t *testing.T) {
		assert := assert.New(t)

		gifAdapterMock := new(gifAdapter.AdapterMock)

		gifAdapterMock.On(
			"SaveContent",
			mock.Anything,
			mock.AnythingOfType("*entities.GIF"),
		).Return(GIFWithURLStub(), nil)

		gifAdapterMock.On(
			"SaveMetadata",
			mock.Anything,
			mock.AnythingOfType("*entities.GIF"),
		).Return(nil, errors.New("unhandled error"))

		gifAdapterMock.On(
			"RemoveContent",
			mock.Anything,
			mock.AnythingOfType("*entities.GIF"),
		).Return(GIFWithURLStub(), nil)

		useCase := MakeGIFUseCase(gifAdapterMock)

		createdGIF, err := useCase.Create(ctx, GIFToBeCreatedStub())

		assert.Nil(createdGIF)
		assert.EqualError(err, "creating metadata error")
		gifAdapterMock.AssertNumberOfCalls(t, "SaveContent", 1)
		gifAdapterMock.AssertNumberOfCalls(t, "SaveMetadata", 1)
		gifAdapterMock.AssertNumberOfCalls(t, "RemoveContent", 1)
	})

	t.Run("Error saving Metadata and removing content", func(t *testing.T) {
		assert := assert.New(t)

		gifAdapterMock := new(gifAdapter.AdapterMock)

		gifAdapterMock.On(
			"SaveContent",
			mock.Anything,
			mock.AnythingOfType("*entities.GIF"),
		).Return(GIFWithURLStub(), nil)

		gifAdapterMock.On(
			"SaveMetadata",
			mock.Anything,
			mock.AnythingOfType("*entities.GIF"),
		).Return(nil, errors.New("unhandled error"))

		gifAdapterMock.On(
			"RemoveContent",
			mock.Anything,
			mock.AnythingOfType("*entities.GIF"),
		).Return(nil, errors.New("unhandled error"))

		useCase := MakeGIFUseCase(gifAdapterMock)

		createdGIF, err := useCase.Create(ctx, GIFToBeCreatedStub())

		assert.Nil(createdGIF)
		assert.EqualError(err, "creating metadata error")
		gifAdapterMock.AssertNumberOfCalls(t, "SaveContent", 1)
		gifAdapterMock.AssertNumberOfCalls(t, "SaveMetadata", 1)
		gifAdapterMock.AssertNumberOfCalls(t, "RemoveContent", 1)
	})
}
