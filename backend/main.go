package main

import (
	"context"
	gifComposite "giphy/challenge/src/infrastructure/composite/gif"
	"giphy/challenge/src/infrastructure/database/mongodb"
	"giphy/challenge/src/infrastructure/database/mongodb/gif/adapters"
	servergraphql "giphy/challenge/src/infrastructure/graphql/server"
	"giphy/challenge/src/infrastructure/http/giphy/adapter"
	"giphy/challenge/src/infrastructure/http/server"
	gifUseCase "giphy/challenge/src/use_cases/gif"
	"log"
	"net/http"
	"time"
)

func main() {
	mongoDBClient, err := mongodb.MakeMongoDBConnection()
	if err != nil {
		log.Fatalf("Error creating MongoDB connection: %s", err.Error())
	}

	gifPort := gifComposite.MakeGIFAdapter(
		adapters.MakeGIFMongoDBAdapter(mongoDBClient),
		adapter.MakeGIFAdapter(&http.Client{Timeout: 60 * time.Second}),
	)

	gifUseCase := gifUseCase.MakeGIFUseCase(gifPort)

	graphQLServer := servergraphql.MakeGraphQLServer(gifUseCase)

	server.MakeAndRunHTTPServer(graphQLServer)

	defer func() {
		ctx := context.Background()

		log.Println("Disconnecting from MongoDB...")
		mongoDBClient.Disconnect(ctx)
	}()
}
