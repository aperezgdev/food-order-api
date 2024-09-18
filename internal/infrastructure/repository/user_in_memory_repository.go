package repository

import (
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/error"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	vo "github.com/aperezgdev/food-order-api/internal/domain/value_object"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/User"
)

type UserInMemoryRepository struct {
	users map[string]entity.User
}

func NewUserInMemoryRepository() repository.UserRepository {
	return &UserInMemoryRepository{
		users: map[string]entity.User{
			"1": {
				Id:        value_object.UserId("1"),
				Name:      value_object.NewUserName("Bob"),
				Email:     value_object.NewUserEmail("bob@bob.com"),
				CreatedOn: vo.NewCreatedOn(),
			},
		},
	}
}

func (uir *UserInMemoryRepository) FindById(id value_object.UserId) (entity.User, error) {
	user, exists := uir.users[string(id)]

	if !exists {
		return user, domain_errors.NotFound
	}

	return user, nil
}

func (uir *UserInMemoryRepository) Save(user entity.User) error {
	uir.users[string(user.Id)] = user

	return nil
}
