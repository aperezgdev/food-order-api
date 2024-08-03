package repository

import . "github.com/aperezgdev/food-order-api/internal/domain/entity"

type DishRepository interface {
	FindAll() (*[]Dish, error)
	Save(*Dish) error
	Update(*Dish) error
}
