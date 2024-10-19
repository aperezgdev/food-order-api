package application

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	"github.com/aperezgdev/food-order-api/internal/domain/shared/result"
)

type OrderFinderAll struct {
	orderRepository repository.OrderRepository
	slog            *slog.Logger
}

func NewOrderFinderAll(
	orderRepository repository.OrderRepository,
	slog *slog.Logger,
) *OrderFinderAll {
	return &OrderFinderAll{orderRepository, slog}
}

func (ofa *OrderFinderAll) Run() *result.Result[[]model.Order] {
	orders, err := ofa.orderRepository.FindAll()
	if err != nil {
		return result.ErrorResult[[]model.Order](domain_errors.Database)
	}

	return result.OkResult(&orders)
}
