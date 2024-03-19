package resolvers

import "giphy/challenge/src/use_cases/gif"

type Resolver struct {
	GIFUseCase gif.GIFUseCaseInterface
}
