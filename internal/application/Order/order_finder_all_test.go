package application

import (
	"log/slog"
	"testing"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	vo "github.com/aperezgdev/food-order-api/internal/domain/value_object"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/Order"
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

	var orders *[]entity.Order
	result.Ok(func(t *[]entity.Order) {
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

	order := entity.Order{
		Id:        value_object.OrderId("2"),
		Status:    value_object.NEW,
		Dishes:    make([]*entity.Dish, 0),
		CreatedOn: vo.NewCreatedOn(),
	}

	orderCreator.Run(order)

	result := orderFinderAll.Run()

	var orders *[]entity.Order
	result.Ok(func(t *[]entity.Order) {
		orders = t
	})

	if len(*orders) != 2 {
		t.Errorf("TestOrderFinderAll - OrderFinderall is not returning all orders")
	}
}
