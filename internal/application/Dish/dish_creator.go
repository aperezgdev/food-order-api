package application

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
)

type DishCreator struct {
	slog           *slog.Logger
	dishRepository repository.DishRepository
}

func NewDishCreator(slog *slog.Logger, dishRepository repository.DishRepository) *DishCreator {
	return &DishCreator{slog, dishRepository}
}

func (dc *DishCreator) Run(dish entity.Dish) error {
	err := dc.dishRepository.Save(dish)
	if err != nil {
		dc.slog.Error("DishCreator.Run", err.Error())
		return err
	}

	return nil
}
