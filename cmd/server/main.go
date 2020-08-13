package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rijine/ads-api/internal/database"
	"github.com/rijine/ads-api/pkg/graph/directive"
	"github.com/rijine/ads-api/pkg/graph/generated"
	"github.com/rijine/ads-api/pkg/graph/resolver"
	"log"
	"net/http"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(database.ErrDatabaseConn)
	}
	config := generated.Config{
		Resolvers: &resolver.Resolver{},
		Directives: generated.DirectiveRoot{
			Auth: directive.Auth,
		},
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	http.Handle("/graphiql", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server")
	}
}
