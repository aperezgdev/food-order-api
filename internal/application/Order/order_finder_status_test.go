package application

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	result "github.com/aperezgdev/food-order-api/internal/domain/shared/result"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/order"
)

// Should return one order with NEW status
func TestOrderFinderStatus(t *testing.T) {
	orderRepository := repository.NewMockOrderRepository()
	orderRepository.On("FindByStatus", mock.Anything).Return([]model.Order{{
		Id:     "1",
		Status: value_object.NEW,
	}}, nil)
	orderFinderStatus := NewOrderFinderStatus(orderRepository, slog.Default())

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

// Shoult return empty list
func TestOrderFinderStatusNotExistsStatus(t *testing.T) {
	orderRepository := repository.NewMockOrderRepository()
	orderRepository.On("FindByStatus", mock.Anything).Return([]model.Order{}, nil)
	orderFinderStatus := NewOrderFinderStatus(orderRepository, &slog.Logger{})

	var res *result.Result[[]model.Order]
	res = orderFinderStatus.Run(value_object.READY)

	var orders []model.Order
	res.Ok(func(t *[]model.Order) {
		orders = *t
	})

	fmt.Println(len(orders))

	if len(orders) != 0 {
		t.Errorf(
			"TestOrderFinderStatus - Result from was searched before created a new ready order was not changed",
		)
	}
}
