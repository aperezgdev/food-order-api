package entity

import (
	vo_shared "github.com/aperezgdev/food-order-api/internal/domain/value_object"
	. "github.com/aperezgdev/food-order-api/internal/domain/value_object/Dish"
)

type Dish struct {
	id          DishId
	name        DishName
	description DishDescription
	price       vo_shared.Price
	createdOn   vo_shared.CreatedOn
}

func NewDish(name DishName, description DishDescription, price vo_shared.Price) *Dish {
	return &Dish{NewDishId(), name, description, price, vo_shared.NewCreatedOn()}
}
