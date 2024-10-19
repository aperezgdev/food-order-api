package repository

import (
	"github.com/aperezgdev/food-order-api/internal/domain/model"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	"github.com/aperezgdev/food-order-api/internal/domain/shared/value_object"
	dish_vo "github.com/aperezgdev/food-order-api/internal/domain/value_object/dish"
)

type DishInMemoryRepository struct {
	dishes map[string]model.Dish
}

func NewDishInMemoryRepository() *DishInMemoryRepository {
	return &DishInMemoryRepository{
		dishes: map[string]model.Dish{
			"1": {
				Id:          dish_vo.DishId("1"),
				Name:        dish_vo.NewDishName("Macarrones"),
				Description: dish_vo.NewDishDescription("Macarrones con tomatico"),
				Price:       value_object.NewPrice(12.2),
				CreatedOn:   value_object.NewCreatedOn(),
			},
			"2": {
				Id:          dish_vo.DishId("2"),
				Name:        dish_vo.NewDishName("Arroz con pollo"),
				Description: dish_vo.NewDishDescription("Arroz con pollo"),
				Price:       value_object.NewPrice(9.2),
				CreatedOn:   value_object.NewCreatedOn(),
			},
		},
	}
}

func (dir *DishInMemoryRepository) FindAll() ([]model.Dish, error) {
	var v []model.Dish

	for _, value := range dir.dishes {
		v = append(v, value)
	}

	return v, nil
}

func (dir *DishInMemoryRepository) Find(id dish_vo.DishId) (model.Dish, error) {
	dish, exists := dir.dishes[string(id)]

	if !exists {
		return dish, domain_errors.NotFound
	}

	return dish, nil
}

func (dir *DishInMemoryRepository) Save(dish model.Dish) error {
	dir.dishes[string(dish.Id)] = dish

	return nil
}

func (dir *DishInMemoryRepository) Update(dish model.Dish) error {
	_, exists := dir.dishes[string(dish.Id)]

	if !exists {
		return domain_errors.NotFound
	}

	dir.dishes[string(dish.Id)] = dish

	return nil
}

func (dir *DishInMemoryRepository) Delete(id dish_vo.DishId) error {
	_, exists := dir.dishes[string(id)]

	if !exists {
		return domain_errors.NotFound
	}

	delete(dir.dishes, string(id))

	return nil
}
