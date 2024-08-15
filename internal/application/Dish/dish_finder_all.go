package application

import (
	"log/slog"

	result "github.com/aperezgdev/food-order-api/internal/domain"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/error"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
)

type DishFinderAll struct {
	dishRepository repository.DishRepository
	slog           *slog.Logger
}

func NewDishFinderAll(dishRepository repository.DishRepository, slog *slog.Logger) *DishFinderAll {
	return &DishFinderAll{dishRepository, slog}
}

func (dfa *DishFinderAll) Run() *result.Result[[]entity.Dish] {
	dfa.slog.Info("DishFinderAll.Run")
	dishes, err := dfa.dishRepository.FindAll()
	if err != nil {
		return result.ErrorResult[[]entity.Dish](domain_errors.Database)
	}

	return result.OkResult(&dishes)
}
