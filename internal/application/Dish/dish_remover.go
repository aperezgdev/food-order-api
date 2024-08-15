package application

import (
	"database/sql"
	"errors"
	"log/slog"

	result "github.com/aperezgdev/food-order-api/internal/domain"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/error"
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

func (df *DishRemover) Run(id value_object.DishId) *result.Result[entity.Dish] {
	err := df.dishRepository.Delete(id)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return result.ErrorResult[entity.Dish](domain_errors.NotFound)
	}

	if err != nil {
		return result.ErrorResult[entity.Dish](domain_errors.Database)
	}

	return result.OkResult(&entity.Dish{})
}
