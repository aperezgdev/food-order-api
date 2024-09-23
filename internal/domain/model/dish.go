package model

import (
	vo_shared "github.com/aperezgdev/food-order-api/internal/domain/shared/value_object"
	. "github.com/aperezgdev/food-order-api/internal/domain/value_object/dish"
)

type Dish struct {
	Id          DishId              `gorm:"type:uuid;default:gen_random_uuid()" binding:"-"`
	Name        DishName            `                                           binding:"required"`
	Description DishDescription     `                                           binding:"required"`
	Price       vo_shared.Price     `                                           binding:"required"`
	CreatedOn   vo_shared.CreatedOn `gorm:"default:current_timestamp"           binding:"-"`
	Orders      []*Order            `gorm:"many2many:orders_dishes;"`
}

func NewDish(name DishName, description DishDescription, price vo_shared.Price) *Dish {
	return &Dish{NewDishId(), name, description, price, vo_shared.NewCreatedOn(), []*Order{}}
}
