package repository

import (
	. "github.com/aperezgdev/food-order-api/internal/domain/entity"
	vo "github.com/aperezgdev/food-order-api/internal/domain/value_object/Order"
)

type OrderRepository interface {
	FindByStatus(*vo.OrderStatus) ([]Order, error)
	Save(*Order) error
	Update(*Order) error
}
