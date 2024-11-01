package application

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	result "github.com/aperezgdev/food-order-api/internal/domain/shared/result"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/user"
)

type UserFinder struct {
	userRepository repository.UserRepository
	slog           *slog.Logger
}

func NewUserFinder(userRepository repository.UserRepository, slog *slog.Logger) *UserFinder {
	return &UserFinder{userRepository, slog}
}

func (uf *UserFinder) Run(id value_object.UserId) *result.Result[model.User] {
	uf.slog.Info("UserFinder.Run - User Id", slog.Any("id", id))
	user, err := uf.userRepository.FindById(id)
	if err != nil {
		return result.ErrorResult[model.User](err)
	}

	return result.OkResult(&user)
}
