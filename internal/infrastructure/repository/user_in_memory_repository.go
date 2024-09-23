package repository

import (
	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	vo "github.com/aperezgdev/food-order-api/internal/domain/shared/value_object"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/user"
)

type UserInMemoryRepository struct {
	users map[string]model.User
}

func NewUserInMemoryRepository() repository.UserRepository {
	return &UserInMemoryRepository{
		users: map[string]model.User{
			"1": {
				Id:        value_object.UserId("1"),
				Name:      value_object.NewUserName("Bob"),
				Email:     value_object.NewUserEmail("bob@bob.com"),
				CreatedOn: vo.NewCreatedOn(),
			},
		},
	}
}

func (uir *UserInMemoryRepository) FindById(id value_object.UserId) (model.User, error) {
	user, exists := uir.users[string(id)]

	if !exists {
		return user, domain_errors.NotFound
	}

	return user, nil
}

func (uir *UserInMemoryRepository) Save(user model.User) error {
	uir.users[string(user.Id)] = user

	return nil
}
