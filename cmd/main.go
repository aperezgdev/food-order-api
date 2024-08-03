package main

import (
	"log"

	"go.uber.org/fx"

	postgres_handler "github.com/aperezgdev/food-order-api/internal/infrastructure/postgres"
)

func main() {
	fx.New(
		fx.Provide(
			log.New,
			postgres_handler.NewPostgresHandler,
		),
	).Run()
}
