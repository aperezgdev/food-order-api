package main

import (
	"log"

	"go.uber.org/fx"

	"github.com/aperezgdev/food-order-api/env"
	postgres_handler "github.com/aperezgdev/food-order-api/internal/infrastructure/postgres"
)

func main() {
	fx.New(
		fx.Provide(
			env.NewEnvApp,
			log.New,
			postgres_handler.NewPostgresHandler,
		),
	).Run()
}
