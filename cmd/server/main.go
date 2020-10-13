package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"github.com/rijine/ads-api/internal/database"
	"github.com/rijine/ads-api/pkg/graph/directive"
	"github.com/rijine/ads-api/pkg/graph/generated"
	"github.com/rijine/ads-api/pkg/graph/resolver"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(helmet.New())

	cfg := limiter.Config{
		// Add more
		Max: 10,
	}
	app.Use(limiter.New(cfg))
	if err := database.Connect(); err != nil {
		log.Fatal(database.ErrDatabaseConn)
	}
	config := generated.Config{
		Resolvers: &resolver.Resolver{},
		Directives: generated.DirectiveRoot{
			Auth:      directive.Auth,
			Maxlength: directive.MaxLength,
			Demo:      directive.Demo,
		},
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	app.Get("/graphiql", adaptor.HTTPHandler(playground.Handler("Ads API", "/graphql")))
	app.Post("/graphql", adaptor.HTTPHandler(srv))

	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Failed to start server")
	}
}
