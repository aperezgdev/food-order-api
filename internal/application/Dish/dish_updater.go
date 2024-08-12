package application

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
)

type DishUpdater struct {
	dishRepository repository.DishRepository
	slog           *slog.Logger
}

func NewDishUpdater(dishRepository repository.DishRepository, slog *slog.Logger) *DishUpdater {
	return &DishUpdater{dishRepository, slog}
}

func (du *DishUpdater) Run(dish entity.Dish) error {
	return du.dishRepository.Update(dish)
}
