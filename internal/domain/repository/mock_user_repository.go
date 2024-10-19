package repository

import (
	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	user_vo "github.com/aperezgdev/food-order-api/internal/domain/value_object/user"
)

type MockUserRepository struct {
	mock.Mock
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

func (m *MockUserRepository) FindById(id user_vo.UserId) (model.User, error) {
	args := m.Called(id)

	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserRepository) Save(user model.User) error {
	args := m.Called(user)

	return args.Error(0)
}
