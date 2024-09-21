package application

import (
	"log/slog"
	"testing"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	vo "github.com/aperezgdev/food-order-api/internal/domain/value_object"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/Dish"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/repository"
)

func newDishCreator() *DishCreator {
	return NewDishCreator(slog.Default(), repository.NewDishInMemoryRepository())
}

// Should create a dish
func TestDishCreator(t *testing.T) {
	dishCreator := newDishCreator()

	dish := entity.Dish{
		Id:          value_object.DishId("3"),
		Name:        value_object.DishName("Fish and chips"),
		Description: value_object.DishDescription("Fish with chips"),
		Price:       vo.Price(10),
		CreatedOn:   vo.NewCreatedOn(),
	}

	result := dishCreator.Run(dish)

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if testError != nil {
		t.Errorf("TestDishCreator - Error has ocurred while trying to create a dish")
	}
}

// Should create a creator and save a user
func TestDishCreatorAndSave(t *testing.T) {
	dishRepository := repository.NewDishInMemoryRepository()
	dishCreator := NewDishCreator(slog.Default(), dishRepository)
	dishFinderAll := NewDishFinderAll(dishRepository, slog.Default())

	var nDishBefore int
	resultBefore := dishFinderAll.Run()
	resultBefore.Ok(func(t *[]entity.Dish) {
		nDishBefore = len(*t)
	})

	dish := entity.Dish{
		Id:          value_object.DishId("3"),
		Name:        value_object.DishName("Fish and chips"),
		Description: value_object.DishDescription("Fish with chips"),
		Price:       vo.Price(10),
		CreatedOn:   vo.NewCreatedOn(),
	}

	dishCreator.Run(dish)

	var nDishAfter int
	resultAfter := dishFinderAll.Run()
	resultAfter.Ok(func(t *[]entity.Dish) {
		nDishAfter = len(*t)
	})

	if nDishAfter != nDishBefore+1 {
		t.Errorf("TestDishCreatorAndSave - DishCreator didnt save the dish")
	}
}
