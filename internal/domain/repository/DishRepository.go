package repository

import (
	. "github.com/aperezgdev/food-order-api/internal/domain/entity"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/Dish"
)

type DishRepository interface {
	FindAll() ([]Dish, error)
	Save(Dish) error
	Update(Dish) error
	Delete(id value_object.DishId) error
}
