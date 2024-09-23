package application

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	"github.com/aperezgdev/food-order-api/internal/domain/shared/result"
)

type DishUpdater struct {
	dishRepository repository.DishRepository
	slog           *slog.Logger
}

func NewDishUpdater(dishRepository repository.DishRepository, slog *slog.Logger) *DishUpdater {
	return &DishUpdater{dishRepository, slog}
}

func (du *DishUpdater) Run(dish model.Dish) *result.Result[model.Dish] {
	err := du.dishRepository.Update(dish)
	if err != nil {
		return result.ErrorResult[model.Dish](domain_errors.Database)
	}

	return result.OkResult(&model.Dish{})
}
