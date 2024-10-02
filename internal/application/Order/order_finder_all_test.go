package application

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
)

// Should return all orders
func TestOrderFinderAll(t *testing.T) {
	orderRepository := repository.NewMockOrderRepository()
	orderRepository.On("FindAll").Return([]model.Order{
		{
			Id: "1",
		},
	}, nil)
	orderFinderAll := NewOrderFinderAll(orderRepository, slog.Default())

	result := orderFinderAll.Run()

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if testError != nil {
		t.Errorf("TestOrderFinderAll - OrderFinderAll return an error")
	}

	var orders *[]model.Order
	result.Ok(func(t *[]model.Order) {
		orders = t
	})

	if len(*orders) != 1 {
		t.Errorf("TestOrderFinderAll - Result should have only one order")
	}
}

// Should return 2 orders
func TestOrderFinderAllAfterCreator(t *testing.T) {
	orderRepository := repository.NewMockOrderRepository()
	orderRepository.On("FindAll").Return([]model.Order{}, domain_errors.Database)
	orderFinderAll := NewOrderFinderAll(orderRepository, &slog.Logger{})

	result := orderFinderAll.Run()

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if !errors.Is(testError, domain_errors.Database) {
		t.Errorf("TestOrderFinderAll - OrderFinderall is not returning all orders")
	}
}
