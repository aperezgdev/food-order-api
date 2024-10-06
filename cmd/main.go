package main

import (
	"net/http"

	"go.uber.org/fx"

	"github.com/aperezgdev/food-order-api/env"
	app_dish "github.com/aperezgdev/food-order-api/internal/application/dish"
	app_order "github.com/aperezgdev/food-order-api/internal/application/order"
	app_user "github.com/aperezgdev/food-order-api/internal/application/user"
	http_server "github.com/aperezgdev/food-order-api/internal/infrastructure/http"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/http/controller"
	route_dish "github.com/aperezgdev/food-order-api/internal/infrastructure/http/route/dish"
	route_order "github.com/aperezgdev/food-order-api/internal/infrastructure/http/route/order"
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
			postgres_handler.NewGormPostgresHandler,
			repository.NewUserPostgresRepository,
			repository.NewDishPostgresRepository,
			repository.NewOrderPostgresRepository,
			http_server.NewHTTPGinServer,
			app_user.NewUserCreator,
			app_user.NewUserFinder,
			app_dish.NewDishCreator,
			app_dish.NewDishFinderAll,
			app_dish.NewDishRemover,
			app_dish.NewDishUpdater,
			app_order.NewOrderCreator,
			app_order.NewOrderFinderAll,
			app_order.NewOrderFinderStatus,
			app_order.NewOrderStatusUpdater,
			controller.NewUserController,
			controller.NewDishController,
			controller.NewOrderController,
			fx.Annotate(
				http_server.NewHTTPRouterGinGonic,
				fx.ParamTags(`group:"routes"`),
			),
			asRoute(route_user.NewUserPostRouteHandler),
			asRoute(route_user.NewUserGetRouteHandler),
			asRoute(route_dish.NewDishPutRouteHandler),
			asRoute(route_dish.NewDishPostRouteHandler),
			asRoute(route_dish.NewDishGetRouteHandler),
			asRoute(route_dish.NewDishDeleteRouteHandler),
			asRoute(route_order.NewOrderGetRouteHandler),
			asRoute(route_order.NewOrderGetStatusHandler),
			asRoute(route_order.NewOrderPathStatusRouteHandler),
			asRoute(route_order.NewOrderPostRouteHandler),
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
