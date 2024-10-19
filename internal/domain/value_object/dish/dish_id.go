package dish_vo

import (
	"github.com/google/uuid"
)

type DishId string

func NewDishId() DishId {
	return DishId(uuid.New().String())
}
