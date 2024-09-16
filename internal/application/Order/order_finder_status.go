package application

import (
	"log/slog"

	result "github.com/aperezgdev/food-order-api/internal/domain"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/error"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/Order"
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

func (ofs *OrderFinderStatus) Run(status value_object.OrderStatus) *result.Result[[]entity.Order] {
	orders, err := ofs.orderRepository.FindByStatus(status)
	if err != nil {
		return result.ErrorResult[[]entity.Order](domain_errors.Database)
	}

	return result.OkResult(&orders)
}
