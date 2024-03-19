package main

import (
	servergraphql "giphy/challenge/src/infrastructure/graphql/server"
	"giphy/challenge/src/infrastructure/http/server"
	"giphy/challenge/src/use_cases/gif"
)

func main() {
	gifUseCase := gif.MakeGIFUseCase(nil)

	graphQLServer := servergraphql.MakeGraphQLServer(gifUseCase)

	server.MakeAndRunHTTPServer(graphQLServer)
}
