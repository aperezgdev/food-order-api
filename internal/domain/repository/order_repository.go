package repository

import (
	"github.com/aperezgdev/food-order-api/internal/domain/model"
	order_vo "github.com/aperezgdev/food-order-api/internal/domain/value_object/order"
)

type OrderRepository interface {
	FindAll() ([]model.Order, error)
	FindByStatus(order_vo.OrderStatus) ([]model.Order, error)
	Save(model.Order) error
	UpdateStatus(order_vo.OrderId, order_vo.OrderStatus) error
}
