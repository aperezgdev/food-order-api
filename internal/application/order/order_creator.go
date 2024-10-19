package application

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	"github.com/aperezgdev/food-order-api/internal/domain/shared/result"
)

type OrderCreator struct {
	orderRepository repository.OrderRepository
	slog            *slog.Logger
}

func NewOrderCreator(orderRepository repository.OrderRepository, slog *slog.Logger) *OrderCreator {
	return &OrderCreator{orderRepository, slog}
}

func (oc *OrderCreator) Run(order model.Order) *result.Result[model.Order] {
	err := oc.orderRepository.Save(order)
	if err != nil {
		return result.ErrorResult[model.Order](domain_errors.Database)
	}

	return result.OkResult(&model.Order{})
}
