package entity

import (
	vo_shared "github.com/aperezgdev/food-order-api/internal/domain/value_object"
	vo_dish "github.com/aperezgdev/food-order-api/internal/domain/value_object/Dish"
	. "github.com/aperezgdev/food-order-api/internal/domain/value_object/Order"
)

type Order struct {
	id        OrderId
	status    OrderStatus
	dishes    []vo_dish.DishId
	createdOn vo_shared.CreatedOn
}

func NewOrder(dishes []vo_dish.DishId) Order {
	return Order{
		id: NewOrderId(), 
		status: NEW, 
		dishes: dishes, 
		createdOn: vo_shared.NewCreatedOn(),
	}
}
