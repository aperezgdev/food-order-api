package application

import (
	"database/sql"
	"errors"
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	"github.com/aperezgdev/food-order-api/internal/domain/shared/result"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/order"
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
) *result.Result[model.Order] {
	err := osu.orderRepository.UpdateStatus(id, status)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return result.ErrorResult[model.Order](domain_errors.NotFound)
	}

	if err != nil {
		return result.ErrorResult[model.Order](domain_errors.Database)
	}

	return result.OkResult(&model.Order{})
}
