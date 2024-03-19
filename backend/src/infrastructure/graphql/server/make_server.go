package server

import (
	"giphy/challenge/src/infrastructure/graphql/resolvers"
	"giphy/challenge/src/infrastructure/graphql/server/config"
	"giphy/challenge/src/use_cases/gif"

	"github.com/99designs/gqlgen/graphql/handler"
)

func MakeGraphQLServer(
	gifUseCase gif.GIFUseCaseInterface,
) *handler.Server {
	resolvers := &resolvers.Resolver{
		GIFUseCase: gifUseCase,
	}

	srv := handler.NewDefaultServer(config.NewExecutableSchema(config.Config{Resolvers: resolvers}))

	return srv
}
