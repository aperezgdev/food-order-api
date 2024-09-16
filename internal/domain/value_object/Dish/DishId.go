package value_object

import (
	"github.com/google/uuid"
)

type DishId uuid.UUID

func NewDishId() DishId {
	return DishId(uuid.New())
}
