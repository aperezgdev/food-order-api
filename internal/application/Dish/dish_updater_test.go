package application

import (
	"log/slog"
	"testing"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	vo "github.com/aperezgdev/food-order-api/internal/domain/value_object"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/Dish"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/repository"
)

func newTestDishUpdater() *DishUpdater {
	return NewDishUpdater(repository.NewDishInMemoryRepository(), slog.Default())
}

// Should run without error
func TestDishUpdaterWithoutError(t *testing.T) {
	dishUpdater := newTestDishUpdater()

	dish := entity.Dish{
		Id:          value_object.DishId("1"),
		Name:        value_object.DishName("Fish and chips"),
		Description: value_object.DishDescription("Fish with chips"),
		Price:       vo.Price(10),
		CreatedOn:   vo.NewCreatedOn(),
	}

	result := dishUpdater.Run(dish)

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if testError != nil {
		t.Errorf("TestDishUpdater - Error has ocurred trying to update dish")
	}
}
