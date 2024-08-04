package main

import (
	"log"
	"net/http"

	"go.uber.org/fx"

	"github.com/aperezgdev/food-order-api/env"
	http_server "github.com/aperezgdev/food-order-api/internal/infrastructure/http"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/http/route"
	postgres_handler "github.com/aperezgdev/food-order-api/internal/infrastructure/postgres"
)

func main() {
	fx.New(
		fx.Provide(
			env.NewEnvApp,
			log.Default,
			postgres_handler.NewPostgresHandler,
			http_server.NewHTTPServer,
			fx.Annotate(
				http_server.NewServerMux,
				fx.ParamTags(`group:"routes"`),
			),
			fx.Annotate(
				route.NewEchoHandler,
				fx.As(new(http_server.Route)),
				fx.ResultTags(`group:"routes"`),
			),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
