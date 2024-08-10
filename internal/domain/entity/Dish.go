package entity

import (
	vo_shared "github.com/aperezgdev/food-order-api/internal/domain/value_object"
	. "github.com/aperezgdev/food-order-api/internal/domain/value_object/Dish"
)

type Dish struct {
	Id          DishId              `db:"id"`
	Name        DishName            `db:"name"`
	Description DishDescription     `db:"description"`
	Price       vo_shared.Price     `db:"price"`
	CreatedOn   vo_shared.CreatedOn `db:"createdOn"`
}

func NewDish(name DishName, description DishDescription, price vo_shared.Price) *Dish {
	return &Dish{NewDishId(), name, description, price, vo_shared.NewCreatedOn()}
}
