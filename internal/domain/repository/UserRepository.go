package repository

import (
	. "github.com/aperezgdev/food-order-api/internal/domain/entity"
	vo "github.com/aperezgdev/food-order-api/internal/domain/value_object/User"
)

type UserRepository interface {
	FindById(id vo.UserId) (User, error)
	Save(user User) error
}
