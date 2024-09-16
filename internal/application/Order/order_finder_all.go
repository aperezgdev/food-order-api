package application

import (
	"log/slog"

	result "github.com/aperezgdev/food-order-api/internal/domain"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/error"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
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

func (ofa *OrderFinderAll) Run() *result.Result[[]entity.Order] {
	orders, err := ofa.orderRepository.FindAll()
	if err != nil {
		return result.ErrorResult[[]entity.Order](domain_errors.Database)
	}

	return result.OkResult(&orders)
}
