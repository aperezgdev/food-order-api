package entity

import (
	vo_shared "github.com/aperezgdev/food-order-api/internal/domain/value_object"
	. "github.com/aperezgdev/food-order-api/internal/domain/value_object/Dish"
)

type Dish struct {
	Id          DishId              `gorm:"type:uuid;default:gen_random_uuid()" binding:"-"`
	Name        DishName            `                                           binding:"required" db:"name"`
	Description DishDescription     `                                           binding:"required" db:"description"`
	Price       vo_shared.Price     `                                           binding:"required" db:"price"`
	CreatedOn   vo_shared.CreatedOn `                                           binding:"-"        db:"createdon"`
	Orders      []*Order            `gorm:"many2many:orders_dishes;"`
}

func NewDish(name DishName, description DishDescription, price vo_shared.Price) *Dish {
	return &Dish{NewDishId(), name, description, price, vo_shared.NewCreatedOn(), []*Order{}}
}
