package adapter

import (
	"context"
	"giphy/challenge/src/domain/entities"
	"log"
)

func (g *gifHTTPAdapter) RemoveContent(ctx context.Context, gif *entities.GIF) (*entities.GIF, error) {
	log.Println("[HTTP > Giphy > SaveContent] not supported yet")

	return gif, nil
}
