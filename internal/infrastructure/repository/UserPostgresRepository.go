package repository

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/User"
	postgres_handler "github.com/aperezgdev/food-order-api/internal/infrastructure/postgres"
)

type UserPostgresRepository struct {
	log                 *slog.Logger
	gormPostgresHandler postgres_handler.GormPostgresHandler
}

func NewUserPostgresRepository(
	log *slog.Logger,
	gormPostgresHandler postgres_handler.GormPostgresHandler,
) repository.UserRepository {
	return &UserPostgresRepository{log, gormPostgresHandler}
}

func (ur *UserPostgresRepository) FindById(id value_object.UserId) (entity.User, error) {
	user := entity.User{}
	ctx := ur.gormPostgresHandler.DB.Find(&user, id)
	if ctx.Error != nil {
		return user, ctx.Error
	}

	return user, nil
}

func (ur *UserPostgresRepository) Save(user entity.User) error {
	ctx := ur.gormPostgresHandler.DB.Create(&user)
	if ctx.Error != nil {
		ur.log.Error("UserPostgresRepository.Save", ctx.Error.Error())
	}

	return ctx.Error
}
