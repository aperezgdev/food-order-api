package entity

import (
	vo_shared "github.com/aperezgdev/food-order-api/internal/domain/value_object"
	vo_dish "github.com/aperezgdev/food-order-api/internal/domain/value_object/Dish"
	. "github.com/aperezgdev/food-order-api/internal/domain/value_object/Order"
)

type Order struct {
	Id        OrderId
	Status    OrderStatus
	Dishes    []vo_dish.DishId
	CreatedOn vo_shared.CreatedOn
}

func NewOrder(dishes []vo_dish.DishId) Order {
	return Order{
		Id:        NewOrderId(),
		Status:    NEW,
		Dishes:    dishes,
		CreatedOn: vo_shared.NewCreatedOn(),
	}
}
