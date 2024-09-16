package application

import (
	"log/slog"

	result "github.com/aperezgdev/food-order-api/internal/domain"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/error"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
)

type OrderCreator struct {
	orderRepository repository.OrderRepository
	slog            *slog.Logger
}

func NewOrderCreator(orderRepository repository.OrderRepository, slog *slog.Logger) *OrderCreator {
	return &OrderCreator{orderRepository, slog}
}

func (oc *OrderCreator) Run(order entity.Order) *result.Result[entity.Order] {
	err := oc.orderRepository.Save(order)
	if err != nil {
		return result.ErrorResult[entity.Order](domain_errors.Database)
	}

	return result.OkResult(&entity.Order{})
}
