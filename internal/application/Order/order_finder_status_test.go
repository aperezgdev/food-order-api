package application

import (
	"log/slog"
	"testing"

	result "github.com/aperezgdev/food-order-api/internal/domain"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	vo "github.com/aperezgdev/food-order-api/internal/domain/value_object"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/Order"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/repository"
)

func newTestOrderFinderStatus() *OrderFinderStatus {
	return NewOrderFinderStatus(repository.NewOrderInMemoryRepository(), slog.Default())
}

// Should return one order with NEW status
func TestOrderFinderStatus(t *testing.T) {
	orderFinderStatus := newTestOrderFinderStatus()

	result := orderFinderStatus.Run(value_object.NEW)

	var orders []entity.Order
	result.Ok(func(t *[]entity.Order) {
		orders = *t
	})

	if len(orders) != 1 {
		t.Errorf("TestOrderFinderStatus - OrderFinderStatus is not returing one order")
	}

	if orders[0].Status != value_object.NEW {
		t.Errorf("TestOrderFinderStatus - Order retrieve dont have the asked status")
	}
}

// Sgould return empty list before create, and should retunr one order after create
func TestOrderFinderStatusAfterCreate(t *testing.T) {
	orderRepository := repository.NewOrderInMemoryRepository()
	orderCreator := NewOrderCreator(orderRepository, &slog.Logger{})
	orderFinderStatus := NewOrderFinderStatus(orderRepository, &slog.Logger{})

	var res *result.Result[[]entity.Order]
	res = orderFinderStatus.Run(value_object.READY)

	var ordersReadyBefore int
	res.Ok(func(t *[]entity.Order) {
		ordersReadyBefore = len(*t)
	})

	order := entity.Order{
		Id:        value_object.OrderId("2"),
		Status:    value_object.READY,
		Dishes:    make([]*entity.Dish, 0),
		CreatedOn: vo.NewCreatedOn(),
	}
	orderCreator.Run(order)

	var ordersReadyAfter int
	res = orderFinderStatus.Run(value_object.READY)
	res.Ok(func(t *[]entity.Order) {
		ordersReadyAfter = len(*t)
	})

	if ordersReadyAfter != ordersReadyBefore+1 {
		t.Errorf(
			"TestOrderFinderStatus - Result from was searched before created a new ready order was not changed",
		)
	}
}
