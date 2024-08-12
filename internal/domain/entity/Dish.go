package entity

import (
	vo_shared "github.com/aperezgdev/food-order-api/internal/domain/value_object"
	. "github.com/aperezgdev/food-order-api/internal/domain/value_object/Dish"
)

type Dish struct {
	Id          DishId              `db:"id"          binding:"-"`
	Name        DishName            `db:"name"        binding:"required"`
	Description DishDescription     `db:"description" binding:"required"`
	Price       vo_shared.Price     `db:"price"       binding:"required"`
	CreatedOn   vo_shared.CreatedOn `db:"createdon"   binding:"-"`
}

func NewDish(name DishName, description DishDescription, price vo_shared.Price) *Dish {
	return &Dish{NewDishId(), name, description, price, vo_shared.NewCreatedOn()}
}
