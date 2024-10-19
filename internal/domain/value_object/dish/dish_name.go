package dish_vo

type DishName string

func NewDishName(name string) DishName {
	return DishName(name)
}

func (d *DishName) Validate() bool {
	return len(*d) > 2
}
