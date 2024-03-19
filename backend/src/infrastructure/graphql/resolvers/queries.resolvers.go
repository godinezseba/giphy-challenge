package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"giphy/challenge/src/infrastructure/graphql/model"
	"giphy/challenge/src/infrastructure/graphql/model/mappers"
	"giphy/challenge/src/infrastructure/graphql/server/config"
)

// Gifs is the resolver for the gifs field.
func (r *queryResolver) Gifs(ctx context.Context, input *model.GIFSearchInput) (*model.GIFPagination, error) {
	results, err := r.GIFUseCase.FindByFilter(ctx, mappers.GIFSearchModelToEntity(input))

	if err != nil {
		return nil, err
	}

	return mappers.GifPaginationEntityToModel(results), nil
}

// Query returns config.QueryResolver implementation.
func (r *Resolver) Query() config.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
