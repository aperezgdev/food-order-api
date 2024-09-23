package application

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	"github.com/aperezgdev/food-order-api/internal/domain/shared/result"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/order"
)

type OrderFinderStatus struct {
	orderRepository repository.OrderRepository
	slog            *slog.Logger
}

func NewOrderFinderStatus(
	orderRepository repository.OrderRepository,
	slog *slog.Logger,
) *OrderFinderStatus {
	return &OrderFinderStatus{orderRepository, slog}
}

func (ofs *OrderFinderStatus) Run(status value_object.OrderStatus) *result.Result[[]model.Order] {
	orders, err := ofs.orderRepository.FindByStatus(status)
	if err != nil {
		return result.ErrorResult[[]model.Order](domain_errors.Database)
	}

	return result.OkResult(&orders)
}
