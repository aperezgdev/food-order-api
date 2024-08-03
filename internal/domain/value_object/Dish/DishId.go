package value_object

import "github.com/google/uuid"

type DishId string

func NewDishId() DishId {
	return DishId(uuid.New().String())
}
