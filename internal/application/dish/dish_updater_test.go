package application

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	vo "github.com/aperezgdev/food-order-api/internal/domain/shared/value_object"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/dish"
)

// Should run without error
func TestDishUpdaterWithoutError(t *testing.T) {
	dishRepository := repository.NewMockDishRepository()
	dishRepository.On("Update", mock.Anything).Return(nil)
	dishUpdater := NewDishUpdater(dishRepository, slog.Default())

	dish := model.Dish{
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

// Repository will return error dish updater should wrap error on return
func TestDishUpdaterDishRepository(t *testing.T) {
	dishRepository := repository.NewMockDishRepository()
	dishRepository.On("Update", mock.Anything).Return(domain_errors.Database)
	dishUpdater := NewDishUpdater(dishRepository, slog.Default())

	dish := model.Dish{
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

	if testError == nil {
		t.Errorf("TestDishUpdater - Error should has ocurred trying to update dish")
	}

	if !errors.Is(testError, domain_errors.Database) {
		t.Errorf("TestDishUpdater - Error is not database error from domain erros")
	}
}
