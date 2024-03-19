package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func MakeAndRunHTTPServer(graphQLServer *handler.Server) {
	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "8080"
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Handle("/playground", playground.Handler("GraphQL playground", "/"))
	router.Handle("/", graphQLServer)

	log.Printf("Connect to http://localhost:%s/playground for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
