package application

import (
	"log/slog"

	result "github.com/aperezgdev/food-order-api/internal/domain"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/error"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
)

type DishUpdater struct {
	dishRepository repository.DishRepository
	slog           *slog.Logger
}

func NewDishUpdater(dishRepository repository.DishRepository, slog *slog.Logger) *DishUpdater {
	return &DishUpdater{dishRepository, slog}
}

func (du *DishUpdater) Run(dish entity.Dish) *result.Result[entity.Dish] {
	err := du.dishRepository.Update(dish)
	if err != nil {
		return result.ErrorResult[entity.Dish](domain_errors.Database)
	}

	return result.OkResult(&entity.Dish{})
}
