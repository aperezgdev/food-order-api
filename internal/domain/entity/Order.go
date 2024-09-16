package entity

import (
	vo_shared "github.com/aperezgdev/food-order-api/internal/domain/value_object"
	. "github.com/aperezgdev/food-order-api/internal/domain/value_object/Order"
)

type Order struct {
	Id        OrderId             `gorm:"type:uuid;default:gen_random_uuid()"`
	Status    OrderStatus         `db:"status"`
	Dishes    []*Dish             `gorm:"many2many:orders_dishes;"`
	CreatedOn vo_shared.CreatedOn `db:"createdOn"`
}

func NewOrder(dishes []*Dish) Order {
	return Order{
		Id:        NewOrderId(),
		Status:    NEW,
		Dishes:    dishes,
		CreatedOn: vo_shared.NewCreatedOn(),
	}
}
