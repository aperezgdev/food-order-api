package application

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
)

type UserCreator struct {
	userRepository repository.UserRepository
	log            *slog.Logger
}

func NewUserCreator(userRepository repository.UserRepository, log *slog.Logger) *UserCreator {
	return &UserCreator{userRepository, log}
}

func (uc *UserCreator) Run(user *entity.User) error {
	err := uc.userRepository.Save(*user)
	if err != nil {
		uc.log.Error("UserCreator.Run", user)
	}

	return err
}
