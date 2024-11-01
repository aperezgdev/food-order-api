package model

import (
	vo_shared "github.com/aperezgdev/food-order-api/internal/domain/shared/value_object"
	. "github.com/aperezgdev/food-order-api/internal/domain/value_object/order"
)

type Order struct {
	Id        OrderId             `gorm:"type:uuid;default:gen_random_uuid()"`
	Status    OrderStatus         `gorm:"default:new"`
	Dishes    []*Dish             `gorm:"many2many:orders_dishes;"`
	CreatedOn vo_shared.CreatedOn `gorm:"default:current_timestamp"           binding:"-"`
}

func NewOrder(dishes []*Dish) Order {
	return Order{
		Id:        NewOrderId(),
		Status:    NEW,
		Dishes:    dishes,
		CreatedOn: vo_shared.NewCreatedOn(),
	}
}
