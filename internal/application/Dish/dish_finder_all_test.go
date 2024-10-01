package application

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
)

// Should find all dish without error
func TestDishFinderAll(t *testing.T) {
	dishRepository := repository.NewMockDishRepository()
	dishRepository.On("FindAll").Return([]model.Dish{
		{},
	}, nil)
	dishFinderAll := NewDishFinderAll(dishRepository, slog.Default())

	result := dishFinderAll.Run()

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if testError != nil {
		t.Errorf("TestDishFinderAll - Error has ocurred while trying to find all dishes")
	}

	var dishes []model.Dish
	result.Ok(func(t *[]model.Dish) {
		dishes = *t
	})

	if len(dishes) == 0 {
		t.Errorf("TestDishFinderAll - Returned list was empty")
	}
}

// Repository will return an error and dish finder all should wrap the error on result
func TestDishFinderAllRepositoryError(t *testing.T) {
	dishRepository := repository.NewMockDishRepository()
	dishRepository.On("FindAll").Return([]model.Dish{}, domain_errors.Database)
	dishFinderAll := NewDishFinderAll(dishRepository, slog.Default())

	result := dishFinderAll.Run()

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if testError == nil {
		t.Errorf("TestDishFinderAll - Error should has ocurred while trying to find all dishes")
	}

	if !errors.Is(testError, domain_errors.Database) {
		t.Errorf("TestDishFinderAll - Error is not an domain error")
	}
}
