package gif

import (
	"context"
	"errors"
	"giphy/challenge/src/domain/entities"
	"log"
)

func (g *GIFUseCase) FindByFilter(
	ctx context.Context,
	query *entities.GIFSearch,
) (*entities.Pagination[entities.GIF], error) {
	log.Println(
		"[Use Case > GIF > Find by Filter] Searching...",
		map[string]any{"search": query.Query, "page": query.Page, "pageSize": query.PageSize},
	)

	result, err := g.GIFPersistenceAdapter.FindByFilter(ctx, query)
	if err != nil {
		log.Println(
			"[Use Case > GIF > Find by Filter] Error storing content",
			map[string]any{"search": query.Query, "page": query.Page, "error": err.Error()},
		)

		return nil, errors.New("finding gifs error")
	}

	return result, nil
}
