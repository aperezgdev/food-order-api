package application

import (
	"log/slog"
	"testing"

	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/dish"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/repository"
)

func newTestDishRemover() *DishRemover {
	return NewDishRemover(repository.NewDishInMemoryRepository(), slog.Default())
}

// Should remove without error
func TestDishRemover(t *testing.T) {
	dishRemover := newTestDishRemover()

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
	dishRemover := newTestDishRemover()

	result := dishRemover.Run(value_object.DishId("931"))

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if testError != domain_errors.NotFound {
		t.Errorf("TestDishRemover - Error should returned cause dish not exists")
	}
}
