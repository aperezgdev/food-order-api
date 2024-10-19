package repository

import (
	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	dish_vo "github.com/aperezgdev/food-order-api/internal/domain/value_object/dish"
)

type MockDishRepository struct {
	mock.Mock
}

func NewMockDishRepository() *MockDishRepository {
	return &MockDishRepository{}
}

func (m *MockDishRepository) FindAll() ([]model.Dish, error) {
	args := m.Called()

	return args.Get(0).([]model.Dish), args.Error(1)
}

func (m *MockDishRepository) Save(dish model.Dish) error {
	args := m.Called(dish)
	return args.Error(0)
}

func (m *MockDishRepository) Find(id dish_vo.DishId) (model.Dish, error) {
	args := m.Called(id)

	return args.Get(0).(model.Dish), args.Error(1)
}

func (m *MockDishRepository) Update(dish model.Dish) error {
	args := m.Called(dish)

	return args.Error(0)
}

func (m *MockDishRepository) Delete(id dish_vo.DishId) error {
	args := m.Called(id)

	return args.Error(0)
}
