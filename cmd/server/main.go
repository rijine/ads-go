package main

import (
	"fmt"
	"log"
	"os"

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
	"github.com/rijine/ads-api/pkg/youtubeapi"
)

var (
	youtubeSrv = youtubeapi.NewYoutubeApiService()
)

func main() {

	fo, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println(youtubeSrv.GetSubscribers("alengcm"))

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "${time}] ${status} - ${latency} ${method} ${path} ${bytesSent} ${query}\n",
		TimeFormat: "2006-01-02T15:04:05-0700",
		TimeZone:   "local",
		Output:     fo,
	}))
	app.Use(compress.New())
	app.Use(helmet.New())

	cfg := limiter.Config{
		// Add more
		Max: 1000000,
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

	// if err := app.Listen(":8080"); err != nil {
	// 	log.Fatal("Failed to start server")
	// }
}
