package repository

import (
	"github.com/aperezgdev/food-order-api/internal/domain/model"
	user_vo "github.com/aperezgdev/food-order-api/internal/domain/value_object/user"
)

type UserRepository interface {
	FindById(id user_vo.UserId) (model.User, error)
	Save(user model.User) error
}
