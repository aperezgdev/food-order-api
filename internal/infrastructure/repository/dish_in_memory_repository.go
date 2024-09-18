package repository

import (
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/error"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object"
	vo_dish "github.com/aperezgdev/food-order-api/internal/domain/value_object/Dish"
)

type DishInMemoryRepository struct {
	dishes map[string]entity.Dish
}

func NewDishInMemoryRepository() *DishInMemoryRepository {
	return &DishInMemoryRepository{
		dishes: map[string]entity.Dish{
			"1": {
				Id:          vo_dish.DishId("1"),
				Name:        vo_dish.NewDishName("Macarrones"),
				Description: vo_dish.NewDishDescription("Macarrones con tomatico"),
				Price:       value_object.NewPrice(12.2),
				CreatedOn:   value_object.NewCreatedOn(),
			},
			"2": {
				Id:          vo_dish.DishId("2"),
				Name:        vo_dish.NewDishName("Arroz con pollo"),
				Description: vo_dish.NewDishDescription("Arroz con pollo"),
				Price:       value_object.NewPrice(9.2),
				CreatedOn:   value_object.NewCreatedOn(),
			},
		},
	}
}

func (dir *DishInMemoryRepository) FindAll() ([]entity.Dish, error) {
	v := make([]entity.Dish, len(dir.dishes))

	for _, value := range dir.dishes {
		v = append(v, value)
	}

	return v, nil
}

func (dir *DishInMemoryRepository) Save(dish entity.Dish) error {
	dir.dishes[string(dish.Id)] = dish

	return nil
}

func (dir *DishInMemoryRepository) Update(dish entity.Dish) error {
	_, exists := dir.dishes[string(dish.Id)]

	if !exists {
		return domain_errors.NotFound
	}

	dir.dishes[string(dish.Id)] = dish

	return nil
}

func (dir *DishInMemoryRepository) Delete(id vo_dish.DishId) error {
	_, exists := dir.dishes[string(id)]

	if !exists {
		return domain_errors.NotFound
	}

	delete(dir.dishes, string(id))

	return nil
}
