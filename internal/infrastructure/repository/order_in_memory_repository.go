package repository

import (
	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	vo "github.com/aperezgdev/food-order-api/internal/domain/shared/value_object"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/order"
)

type OrderInMemoryRepository struct {
	orders map[string]model.Order
}

func NewOrderInMemoryRepository() repository.OrderRepository {
	return &OrderInMemoryRepository{
		orders: map[string]model.Order{
			"1": {
				Id:        value_object.OrderId("1"),
				Status:    value_object.NEW,
				Dishes:    make([]*model.Dish, 0),
				CreatedOn: vo.NewCreatedOn(),
			},
		},
	}
}

func (oir *OrderInMemoryRepository) FindAll() ([]model.Order, error) {
	var v []model.Order

	for _, value := range oir.orders {
		v = append(v, value)
	}

	return v, nil
}

func (oir *OrderInMemoryRepository) FindByStatus(
	status value_object.OrderStatus,
) ([]model.Order, error) {
	var v []model.Order

	for _, value := range oir.orders {
		if value.Status == status {
			v = append(v, value)
		}
	}

	return v, nil
}

func (oir *OrderInMemoryRepository) Save(order model.Order) error {
	oir.orders[string(order.Id)] = order

	return nil
}

func (oir *OrderInMemoryRepository) UpdateStatus(
	id value_object.OrderId,
	status value_object.OrderStatus,
) error {
	order, exists := oir.orders[string(id)]

	if !exists {
		return nil
	}

	order.Status = status

	oir.orders[string(id)] = order

	return nil
}
