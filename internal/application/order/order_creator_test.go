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
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/order"
)

// Should create a valid order
func TestOrderCreatorRepositoryError(t *testing.T) {
	orderRepository := repository.NewMockOrderRepository()
	orderRepository.On("Save", mock.Anything).Return(domain_errors.Database)
	orderCreator := NewOrderCreator(orderRepository, slog.Default())

	order := model.Order{
		Id:        value_object.OrderId("2"),
		Status:    value_object.NEW,
		Dishes:    make([]*model.Dish, 0),
		CreatedOn: vo.NewCreatedOn(),
	}

	var erro error
	result := orderCreator.Run(order)
	result.Error(func(err error) {
		erro = err
	})

	if erro == nil || !errors.Is(erro, domain_errors.Database) {
		t.Errorf("TestOrderCreator - OrderCreator is returning an error")
	}
}

// Should create and save a valid order
func TestOrderCreator(t *testing.T) {
	orderRepository := repository.NewMockOrderRepository()
	orderRepository.On("Save", mock.Anything).Return(nil)
	orderCreator := NewOrderCreator(orderRepository, &slog.Logger{})

	order := model.Order{
		Id:        value_object.OrderId("3"),
		Status:    value_object.WORKING_ON,
		Dishes:    make([]*model.Dish, 0),
		CreatedOn: vo.NewCreatedOn(),
	}

	result := orderCreator.Run(order)

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if testError != nil {
		t.Errorf("TestOrderCreator - Error has ocurred trying to create order")
	}
}
