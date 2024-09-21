package application

import (
	"log/slog"
	"testing"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	vo "github.com/aperezgdev/food-order-api/internal/domain/value_object"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/Dish"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/repository"
)

func newTestDishFinderAll() *DishFinderAll {
	return NewDishFinderAll(repository.NewDishInMemoryRepository(), slog.Default())
}

// Should find all dish without error
func TestDishFinderAll(t *testing.T) {
	userFinderAll := newTestDishFinderAll()

	result := userFinderAll.Run()

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if testError != nil {
		t.Errorf("TestDishFinderAll - Error has ocurred while trying to find all dishes")
	}

	var dishes []entity.Dish
	result.Ok(func(t *[]entity.Dish) {
		dishes = *t
	})

	if len(dishes) == 0 {
		t.Errorf("TestDishFinderAll - Returned list was empty")
	}
}

// Should retrieve one more dish after insert
func TestDishFinderAllAfterSave(t *testing.T) {
	dishRepository := repository.NewDishInMemoryRepository()
	dishFinderAll := NewDishFinderAll(dishRepository, slog.Default())
	dishCreator := NewDishCreator(slog.Default(), dishRepository)

	var nUsersBefore int
	dishFinderAll.Run().Ok(func(t *[]entity.Dish) {
		nUsersBefore = len(*t)
	})

	dish := entity.Dish{
		Id:          value_object.DishId("3"),
		Name:        value_object.DishName("Fish and chips"),
		Description: value_object.DishDescription("Fish with chips"),
		Price:       vo.Price(10),
		CreatedOn:   vo.NewCreatedOn(),
	}

	dishCreator.Run(dish)

	var nUsersAfter int
	dishFinderAll.Run().Ok(func(t *[]entity.Dish) {
		nUsersAfter = len(*t)
	})

	if nUsersAfter != nUsersBefore+1 {
		t.Errorf("TestDishFinderAllAfterSave - Didnt find all dishes")
	}
}
