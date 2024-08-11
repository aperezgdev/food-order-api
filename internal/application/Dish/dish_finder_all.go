package application

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
)

type DishFinderAll struct {
	dishRepository repository.DishRepository
	slog           *slog.Logger
}

func NewDishFinderAll(dishRepository repository.DishRepository, slog *slog.Logger) *DishFinderAll {
	return &DishFinderAll{dishRepository, slog}
}

func (dfa *DishFinderAll) Run() ([]entity.Dish, error) {
	return dfa.dishRepository.FindAll()
}
