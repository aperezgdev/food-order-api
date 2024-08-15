package application

import (
	"log/slog"

	result "github.com/aperezgdev/food-order-api/internal/domain"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	errors "github.com/aperezgdev/food-order-api/internal/domain/error"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
)

type UserCreator struct {
	userRepository repository.UserRepository
	log            *slog.Logger
}

func NewUserCreator(userRepository repository.UserRepository, log *slog.Logger) *UserCreator {
	return &UserCreator{userRepository, log}
}

func (uc *UserCreator) Run(user *entity.User) *result.Result[entity.User] {
	uc.log.Info("UserCreator.Run ", slog.Any("user", user))
	err := uc.userRepository.Save(*user)
	if err != nil {
		return result.ErrorResult[entity.User](errors.Database)
	}

	return result.OkResult(&entity.User{})
}
