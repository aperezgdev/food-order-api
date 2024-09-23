package application

import (
	"log/slog"
	"testing"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	vo "github.com/aperezgdev/food-order-api/internal/domain/shared/value_object"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/order"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/repository"
)

func newTestOrderFinderAll() *OrderFinderAll {
	return NewOrderFinderAll(repository.NewOrderInMemoryRepository(), slog.Default())
}

// Should return all orders
func TestOrderFinderAll(t *testing.T) {
	orderFinderAll := newTestOrderFinderAll()

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
	orderRepository := repository.NewOrderInMemoryRepository()
	orderCreator := NewOrderCreator(orderRepository, &slog.Logger{})
	orderFinderAll := NewOrderFinderAll(orderRepository, &slog.Logger{})

	order := model.Order{
		Id:        value_object.OrderId("2"),
		Status:    value_object.NEW,
		Dishes:    make([]*model.Dish, 0),
		CreatedOn: vo.NewCreatedOn(),
	}

	orderCreator.Run(order)

	result := orderFinderAll.Run()

	var orders *[]model.Order
	result.Ok(func(t *[]model.Order) {
		orders = t
	})

	if len(*orders) != 2 {
		t.Errorf("TestOrderFinderAll - OrderFinderall is not returning all orders")
	}
}
