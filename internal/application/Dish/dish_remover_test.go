package application

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/dish"
)

// Should remove without error
func TestDishRemover(t *testing.T) {
	dishRepository := repository.NewMockDishRepository()
	dishRepository.On("Delete", mock.Anything).Return(nil)
	dishRemover := NewDishRemover(dishRepository, slog.Default())

	result := dishRemover.Run(value_object.DishId("1"))

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if testError != nil {
		t.Errorf("TestDishRemover - Dish was not removed")
	}
}

// Should return NOT_FOUND error trying to remove a not exist dish
func TestDishNotExistRemover(t *testing.T) {
	dishRepository := repository.NewMockDishRepository()
	dishRepository.On("Delete", mock.Anything).Return(domain_errors.NotFound)
	dishRemover := NewDishRemover(dishRepository, slog.Default())

	result := dishRemover.Run(value_object.DishId("931"))

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if testError != domain_errors.NotFound {
		t.Errorf("TestDishRemover - Error should returned cause dish not exists")
	}
}
