package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rijine/ads-api/internal/database"
	"github.com/rijine/ads-api/pkg/graph/generated"
	"github.com/rijine/ads-api/pkg/graph/resolver"
	"log"
	"net/http"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(database.ErrDatabaseConn)
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/graphiql", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server")
	}
}
