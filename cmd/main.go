package main

import (
	"net/http"

	"go.uber.org/fx"

	"github.com/aperezgdev/food-order-api/env"
	app_dish "github.com/aperezgdev/food-order-api/internal/application/Dish"
	app_user "github.com/aperezgdev/food-order-api/internal/application/User"
	http_server "github.com/aperezgdev/food-order-api/internal/infrastructure/http"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/http/controller"
	route_dish "github.com/aperezgdev/food-order-api/internal/infrastructure/http/route/dish"
	route_user "github.com/aperezgdev/food-order-api/internal/infrastructure/http/route/user"
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
			repository.NewDishPostgresRepository,
			http_server.NewHTTPGinServer,
			app_user.NewUserCreator,
			app_user.NewUserFinder,
			app_dish.NewDishCreator,
			controller.NewUserController,
			controller.NewDishController,
			fx.Annotate(
				http_server.NewHTTPRouterGinGonic,
				fx.ParamTags(`group:"routes"`),
			),
			asRoute(route_user.NewUserPostRouteHandler),
			asRoute(route_user.NewUserGetRouteHandler),
			asRoute(route_dish.NewDishPostRouteHandler),
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
