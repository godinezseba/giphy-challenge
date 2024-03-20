package gif

import (
	"context"
	"log"
	"sync"
)

func (g *GIFUseCase) Populate(ctx context.Context) {
	searchs := []string{"cats", "dogs", "penguins", "ferret", "axolotl"}
	maxNumberPerResult := 5

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(searchs))

	for _, searchString := range searchs {
		go func(search string) {
			defer waitGroup.Done()
			log.Println("[Use Case > GIF > Populate] Searching with: ", search)

			gifs, err := g.GIFPersistenceAdapter.FindBySearch(ctx, search, maxNumberPerResult)
			if err != nil {
				log.Println(
					"[Use Case > GIF > Populate] Unhandled error searching",
					map[string]string{"search": search, "error": err.Error()},
				)

				return
			}

			if _, err := g.GIFPersistenceAdapter.SaveMultipleMetadata(ctx, gifs); err != nil {
				log.Println(
					"[Use Case > GIF > Populate] Unhandled error populating DB",
					map[string]string{"search": search, "error": err.Error()},
				)

				return
			}

		}(searchString)
	}

	waitGroup.Wait()
}
