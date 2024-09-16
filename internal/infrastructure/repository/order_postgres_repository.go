package repository

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/Order"
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

func (opr *OrderPostgresRepository) FindAll() ([]entity.Order, error) {
	orders := []entity.Order{}
	result := opr.gormPostgresHandler.DB.Find(&orders)

	return orders, result.Error
}

func (opr *OrderPostgresRepository) UpdateStatus(id value_object.OrderId, status value_object.OrderStatus) error {
	result := opr.gormPostgresHandler.DB.Save(&entity.Order{Id: id, Status: status})

	return result.Error
}

func (opr *OrderPostgresRepository) FindByStatus(status value_object.OrderStatus) ([]entity.Order, error) {
	orders := []entity.Order{}
	result := opr.gormPostgresHandler.DB.Where("status = ?", status).Find(&orders)

	return orders, result.Error
}

func (opr *OrderPostgresRepository) Save(order entity.Order) error {
	result := opr.gormPostgresHandler.DB.Save(&order)
	return result.Error
}
