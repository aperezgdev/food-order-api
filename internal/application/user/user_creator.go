package application

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	result "github.com/aperezgdev/food-order-api/internal/domain/shared/result"
)

type UserCreator struct {
	userRepository repository.UserRepository
	log            *slog.Logger
}

func NewUserCreator(userRepository repository.UserRepository, log *slog.Logger) *UserCreator {
	return &UserCreator{userRepository, log}
}

func (uc *UserCreator) Run(user *model.User) *result.Result[model.User] {
	uc.log.Info("UserCreator.Run ", slog.Any("user", user))
	err := uc.userRepository.Save(*user)
	if err != nil {
		return result.ErrorResult[model.User](domain_errors.Database)
	}

	return result.OkResult(&model.User{})
}
