package application

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/User"
)

type UserFinder struct {
	userRepository repository.UserRepository
	slog           *slog.Logger
}

func NewUserFinder(userRepository repository.UserRepository, slog *slog.Logger) *UserFinder {
	return &UserFinder{userRepository, slog}
}

func (uf *UserFinder) Run(id value_object.UserId) (entity.User, error) {
	uf.slog.Info("UserFinder.Run - User Id", id)
	user, err := uf.userRepository.FindById(id)
	if err != nil {
		uf.slog.Error("UserFinder.Run - Error", err.Error())
		return entity.User{}, err
	}

	return user, nil
}
