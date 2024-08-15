package application

import (
	"database/sql"
	"errors"
	"log/slog"

	result "github.com/aperezgdev/food-order-api/internal/domain"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/error"
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

func (uf *UserFinder) Run(id value_object.UserId) *result.Result[entity.User] {
	uf.slog.Info("UserFinder.Run - User Id", slog.Any("id", id))
	user, err := uf.userRepository.FindById(id)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return result.ErrorResult[entity.User](domain_errors.NotFound)
	}

	if err != nil {
		return result.ErrorResult[entity.User](domain_errors.Database)
	}

	return result.OkResult(&user)
}
