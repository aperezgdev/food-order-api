package application

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	"github.com/aperezgdev/food-order-api/internal/domain/shared/value_object"
	"github.com/aperezgdev/food-order-api/internal/domain/value_object/dish"
)

// Should create a dish
func TestDishCreator(t *testing.T) {
	dishRepository := repository.NewMockDishRepository()
	dishRepository.On("Save", mock.Anything).Return(nil)
	dishCreator := NewDishCreator(slog.Default(), dishRepository)

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

// Repository return an error and dish creator should return an result with error
func TestDishCreatorRepositoryError(t *testing.T) {
	dishRepository := repository.NewMockDishRepository()
	dishRepository.On("Save", mock.Anything).Return(domain_errors.Database)
	dishCreator := NewDishCreator(slog.Default(), dishRepository)

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

	if testError == nil {
		t.Errorf("TestDishCreator - Error should has ocurred")
	}
}
