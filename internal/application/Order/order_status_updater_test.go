package application

import (
	"log/slog"
	"testing"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/Order"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/repository"
)

func newTestOrderStatusUpdater() *OrderStatusUpdater {
	return NewOrderStatusUpdater(repository.NewOrderInMemoryRepository(), &slog.Logger{})
}

// Should update without error
func TestOrderStatusUpdater(t *testing.T) {
	orderStatusUpdater := newTestOrderStatusUpdater()

	result := orderStatusUpdater.Run(value_object.OrderId("1"), value_object.READY)

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if testError != nil {
		t.Errorf("TestOrderStatusUpdater - Error trying to update status from order")
	}
}

// Should update and confirm changes
func TestOrderStatusUpdaterCommmitChanges(t *testing.T) {
	orderRepository := repository.NewOrderInMemoryRepository()
	orderStatusUpdater := NewOrderStatusUpdater(orderRepository, &slog.Logger{})
	orderFinderStatus := NewOrderFinderStatus(orderRepository, &slog.Logger{})

	var nReadyOrderBefore int
	resultBefore := orderFinderStatus.Run(value_object.READY)
	resultBefore.Ok(func(t *[]entity.Order) {
		nReadyOrderBefore = len(*t)
	})

	orderStatusUpdater.Run(value_object.OrderId("1"), value_object.READY)

	var nReadyOrderAfter int
	resultAfter := orderFinderStatus.Run(value_object.READY)
	resultAfter.Ok(func(t *[]entity.Order) {
		nReadyOrderAfter = len(*t)
	})

	if nReadyOrderAfter != nReadyOrderBefore+1 {
		t.Errorf("TestOrderStatusUpdaterCommitChanges - Order status is not updating")
	}
}
