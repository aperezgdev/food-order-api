package application

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	"github.com/aperezgdev/food-order-api/internal/domain/shared/result"
)

type DishFinderAll struct {
	dishRepository repository.DishRepository
	slog           *slog.Logger
}

func NewDishFinderAll(dishRepository repository.DishRepository, slog *slog.Logger) *DishFinderAll {
	return &DishFinderAll{dishRepository, slog}
}

func (dfa *DishFinderAll) Run() *result.Result[[]model.Dish] {
	dfa.slog.Info("DishFinderAll.Run")
	dishes, err := dfa.dishRepository.FindAll()
	if err != nil {
		return result.ErrorResult[[]model.Dish](domain_errors.Database)
	}

	return result.OkResult(&dishes)
}
