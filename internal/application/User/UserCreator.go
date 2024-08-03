package application

import (
	"log"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
)

type UserCreator struct {
	userRepository repository.UserRepository
	log            *log.Logger
}

func NewUserCreator(userRepository repository.UserRepository, log *log.Logger) UserCreator {
	return UserCreator{userRepository, log}
}

func (uc *UserCreator) Run(user entity.User) error {
	err := uc.userRepository.Save(user)
	if err != nil {
		uc.log.Panic("UserCreator.Run", user)
	}

	return err
}
