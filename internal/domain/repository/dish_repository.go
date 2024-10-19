package repository

import (
	"github.com/aperezgdev/food-order-api/internal/domain/model"
	dish_vo "github.com/aperezgdev/food-order-api/internal/domain/value_object/dish"
)

type DishRepository interface {
	FindAll() ([]model.Dish, error)
	Find(dish_vo.DishId) (model.Dish, error)
	Save(model.Dish) error
	Update(model.Dish) error
	Delete(dish_vo.DishId) error
}
