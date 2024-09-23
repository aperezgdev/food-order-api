package dish_vo

type DishDescription string

func NewDishDescription(description string) DishDescription {
	return DishDescription(description)
}

func (d *DishDescription) Validate() bool {
	return len(*d) > 2
}
