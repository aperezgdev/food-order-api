package application

import (
	"log/slog"

	result "github.com/aperezgdev/food-order-api/internal/domain"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	"github.com/aperezgdev/food-order-api/internal/domain/error"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
)

type DishCreator struct {
	slog           *slog.Logger
	dishRepository repository.DishRepository
}

func NewDishCreator(slog *slog.Logger, dishRepository repository.DishRepository) *DishCreator {
	return &DishCreator{slog, dishRepository}
}

func (dc *DishCreator) Run(dish entity.Dish) *result.Result[entity.Dish] {
	dc.slog.Info("DishCreator.Run ", slog.Any("dish", dish))
	err := dc.dishRepository.Save(dish)
	if err != nil {
		return result.ErrorResult[entity.Dish](domain_errors.Database)
	}

	return result.OkResult(&entity.Dish{})
}
