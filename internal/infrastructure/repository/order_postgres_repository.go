package repository

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/order"
	postgres_handler "github.com/aperezgdev/food-order-api/internal/infrastructure/postgres"
)

type OrderPostgresRepository struct {
	slog                *slog.Logger
	gormPostgresHandler postgres_handler.GormPostgresHandler
}

func NewOrderPostgresRepository(
	slog *slog.Logger,
	gormPostgresHandler postgres_handler.GormPostgresHandler,
) repository.OrderRepository {
	return &OrderPostgresRepository{slog, gormPostgresHandler}
}

func (opr *OrderPostgresRepository) FindAll() ([]model.Order, error) {
	orders := []model.Order{}
	result := opr.gormPostgresHandler.DB.Find(&orders)

	return orders, result.Error
}

func (opr *OrderPostgresRepository) UpdateStatus(
	id value_object.OrderId,
	status value_object.OrderStatus,
) error {
	result := opr.gormPostgresHandler.DB.Model(&model.Order{}).
		Where("id = ?", id).
		Update("status", status)

	return result.Error
}

func (opr *OrderPostgresRepository) FindByStatus(
	status value_object.OrderStatus,
) ([]model.Order, error) {
	orders := []model.Order{}
	result := opr.gormPostgresHandler.DB.Where("status = ?", status).Find(&orders)

	return orders, result.Error
}

func (opr *OrderPostgresRepository) Save(order model.Order) error {
	result := opr.gormPostgresHandler.DB.Save(&order)
	return result.Error
}
