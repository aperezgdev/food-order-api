package application

import (
	"errors"
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/order"
)

var (
	orderRepository    = repository.NewMockOrderRepository()
	orderStatusUpdater = NewOrderStatusUpdater(orderRepository, slog.Default())
)

func repositoryShouldReturn(result error) {
	orderRepository.On("UpdateStatus", mock.Anything, mock.Anything).Return(result).Once()
}

// Should update without error
func TestOrderStatusUpdater(t *testing.T) {
	repositoryShouldReturn(nil)

	result := orderStatusUpdater.Run(value_object.OrderId("1"), value_object.READY)

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if testError != nil {
		t.Errorf("TestOrderStatusUpdater - Error trying to update status from order")
	}
}

// Should return error
func TestOrderStatusUpdaterRepositoryError(t *testing.T) {
	repositoryShouldReturn(domain_errors.Database)

	result := orderStatusUpdater.Run(value_object.OrderId("1"), value_object.READY)

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	fmt.Println(testError)

	if !errors.Is(testError, domain_errors.Database) {
		t.Errorf("TestOrderStatusUpdaterRepositoryError - Error didnt ocurred")
	}
}
