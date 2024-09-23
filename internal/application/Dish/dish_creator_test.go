package application

import (
	"log/slog"
	"testing"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/shared/value_object"
	"github.com/aperezgdev/food-order-api/internal/domain/value_object/dish"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/repository"
)

func newDishCreator() *DishCreator {
	return NewDishCreator(slog.Default(), repository.NewDishInMemoryRepository())
}

// Should create a dish
func TestDishCreator(t *testing.T) {
	dishCreator := newDishCreator()

	dish := model.Dish{
		Id:          dish_vo.DishId("3"),
		Name:        dish_vo.DishName("Fish and chips"),
		Description: dish_vo.DishDescription("Fish with chips"),
		Price:       value_object.Price(10),
		CreatedOn:   value_object.NewCreatedOn(),
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
	resultBefore.Ok(func(t *[]model.Dish) {
		nDishBefore = len(*t)
	})

	dish := model.Dish{
		Id:          dish_vo.DishId("3"),
		Name:        dish_vo.DishName("Fish and chips"),
		Description: dish_vo.DishDescription("Fish with chips"),
		Price:       value_object.Price(10),
		CreatedOn:   value_object.NewCreatedOn(),
	}

	dishCreator.Run(dish)

	var nDishAfter int
	resultAfter := dishFinderAll.Run()
	resultAfter.Ok(func(t *[]model.Dish) {
		nDishAfter = len(*t)
	})

	if nDishAfter != nDishBefore+1 {
		t.Errorf("TestDishCreatorAndSave - DishCreator didnt save the dish")
	}
}
