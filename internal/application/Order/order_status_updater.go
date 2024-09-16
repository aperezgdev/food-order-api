package application

import (
	"database/sql"
	"errors"
	"log/slog"

	result "github.com/aperezgdev/food-order-api/internal/domain"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/error"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/Order"
)

type OrderStatusUpdater struct {
	orderRepository repository.OrderRepository
	slog            *slog.Logger
}

func NewOrderStatusUpdater(
	orderRepository repository.OrderRepository,
	slog *slog.Logger,
) *OrderStatusUpdater {
	return &OrderStatusUpdater{orderRepository, slog}
}

func (osu *OrderStatusUpdater) Run(
	id value_object.OrderId,
	status value_object.OrderStatus,
) *result.Result[entity.Order] {
	err := osu.orderRepository.UpdateStatus(id, status)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return result.ErrorResult[entity.Order](domain_errors.NotFound)
	}

	if err != nil {
		return result.ErrorResult[entity.Order](domain_errors.Database)
	}

	return result.OkResult(&entity.Order{})
}
