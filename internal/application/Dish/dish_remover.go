package application

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/Dish"
)

type DishRemover struct {
	dishRepository repository.DishRepository
	slog           *slog.Logger
}

func NewDishRemover(dishRepository repository.DishRepository, slog *slog.Logger) *DishRemover {
	return &DishRemover{dishRepository, slog}
}

func (df *DishRemover) Run(id value_object.DishId) error {
	return df.dishRepository.Delete(id)
}
