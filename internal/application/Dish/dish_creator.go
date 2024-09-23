package application

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	"github.com/aperezgdev/food-order-api/internal/domain/shared/result"
)

type DishCreator struct {
	slog           *slog.Logger
	dishRepository repository.DishRepository
}

func NewDishCreator(slog *slog.Logger, dishRepository repository.DishRepository) *DishCreator {
	return &DishCreator{slog, dishRepository}
}

func (dc *DishCreator) Run(dish model.Dish) *result.Result[model.Dish] {
	dc.slog.Info("DishCreator.Run ", slog.Any("dish", dish))
	err := dc.dishRepository.Save(dish)
	if err != nil {
		return result.ErrorResult[model.Dish](domain_errors.Database)
	}

	return result.OkResult(&model.Dish{})
}
