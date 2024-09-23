package application

import (
	"log/slog"
	"testing"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	result "github.com/aperezgdev/food-order-api/internal/domain/shared/result"
	vo "github.com/aperezgdev/food-order-api/internal/domain/shared/value_object"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/order"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/repository"
)

func newTestOrderFinderStatus() *OrderFinderStatus {
	return NewOrderFinderStatus(repository.NewOrderInMemoryRepository(), slog.Default())
}

// Should return one order with NEW status
func TestOrderFinderStatus(t *testing.T) {
	orderFinderStatus := newTestOrderFinderStatus()

	result := orderFinderStatus.Run(value_object.NEW)

	var orders []model.Order
	result.Ok(func(t *[]model.Order) {
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

	var res *result.Result[[]model.Order]
	res = orderFinderStatus.Run(value_object.READY)

	var ordersReadyBefore int
	res.Ok(func(t *[]model.Order) {
		ordersReadyBefore = len(*t)
	})

	order := model.Order{
		Id:        value_object.OrderId("2"),
		Status:    value_object.READY,
		Dishes:    make([]*model.Dish, 0),
		CreatedOn: vo.NewCreatedOn(),
	}
	orderCreator.Run(order)

	var ordersReadyAfter int
	res = orderFinderStatus.Run(value_object.READY)
	res.Ok(func(t *[]model.Order) {
		ordersReadyAfter = len(*t)
	})

	if ordersReadyAfter != ordersReadyBefore+1 {
		t.Errorf(
			"TestOrderFinderStatus - Result from was searched before created a new ready order was not changed",
		)
	}
}
