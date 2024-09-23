package application

import (
	"errors"
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	"github.com/aperezgdev/food-order-api/internal/domain/shared/result"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/dish"
)

type DishRemover struct {
	dishRepository repository.DishRepository
	slog           *slog.Logger
}

func NewDishRemover(dishRepository repository.DishRepository, slog *slog.Logger) *DishRemover {
	return &DishRemover{dishRepository, slog}
}

func (df *DishRemover) Run(id value_object.DishId) *result.Result[model.Dish] {
	err := df.dishRepository.Delete(id)

	if err != nil && errors.Is(err, domain_errors.NotFound) {
		return result.ErrorResult[model.Dish](domain_errors.NotFound)
	}

	if err != nil {
		return result.ErrorResult[model.Dish](domain_errors.Database)
	}

	return result.OkResult(&model.Dish{})
}
