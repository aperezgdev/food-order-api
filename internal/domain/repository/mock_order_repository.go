package repository

import (
	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	order_vo "github.com/aperezgdev/food-order-api/internal/domain/value_object/order"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) FindAll() ([]model.Order, error) {
	args := m.Called()

	return args.Get(0).([]model.Order), args.Error(1)
}

func (m *MockOrderRepository) FindByStatus(status order_vo.OrderStatus) ([]model.Order, error) {
	args := m.Called(status)

	return args.Get(0).([]model.Order), args.Error(1)
}

func (m *MockOrderRepository) Save(order model.Order) error {
	args := m.Called(order)

	return args.Error(0)
}

func (m *MockOrderRepository) UpdateStatus(id order_vo.OrderId, status order_vo.OrderStatus) error {
	args := m.Called(id, status)

	return args.Error(0)
}
