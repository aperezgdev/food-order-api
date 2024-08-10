package main

import (
	"net/http"

	"go.uber.org/fx"

	"github.com/aperezgdev/food-order-api/env"
	application "github.com/aperezgdev/food-order-api/internal/application/User"
	http_server "github.com/aperezgdev/food-order-api/internal/infrastructure/http"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/http/controller"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/http/route"
	logger "github.com/aperezgdev/food-order-api/internal/infrastructure/log"
	postgres_handler "github.com/aperezgdev/food-order-api/internal/infrastructure/postgres"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/repository"
)

func main() {
	fx.New(
		fx.Provide(
			env.NewEnvApp,
			logger.NewLogger,
			postgres_handler.NewPostgresHandler,
			repository.NewUserPostgresRepository,
			http_server.NewHTTPGinServer,
			application.NewUserCreator,
			controller.NewUserController,
			application.NewUserFinder,
			fx.Annotate(
				http_server.NewHTTPRouterGinGonic,
				fx.ParamTags(`group:"routes"`),
			),
			asRoute(route.NewUserPostRouteHandler),
			asRoute(route.NewUserGetRouteHandler),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func asRoute(route any) interface{} {
	return fx.Annotate(
		route,
		fx.As(new(http_server.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
