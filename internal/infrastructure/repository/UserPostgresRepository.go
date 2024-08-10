package repository

import (
	"fmt"
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/User"
	postgres_handler "github.com/aperezgdev/food-order-api/internal/infrastructure/postgres"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/postgres/queries"
)

type UserPostgresRepository struct {
	log             *slog.Logger
	postgresHandler postgres_handler.PostgresHandler
}

func NewUserPostgresRepository(
	log *slog.Logger,
	postgresHandler postgres_handler.PostgresHandler,
) repository.UserRepository {
	return &UserPostgresRepository{log, postgresHandler}
}

func (ur *UserPostgresRepository) FindById(id value_object.UserId) (entity.User, error) {
	user := entity.User{}
	err := ur.postgresHandler.DB.Get(&user, queries.UserFinder, id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserPostgresRepository) Save(user entity.User) error {
	_, err := ur.postgresHandler.DB.NamedExec(queries.UserCreate, &user)
	if err != nil {
		fmt.Println(err)
		ur.log.Error("UserPostgresRepository.Save", err.Error())
	}

	return err
}
