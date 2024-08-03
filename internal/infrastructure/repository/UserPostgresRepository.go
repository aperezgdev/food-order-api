package repository

import (
	"log"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/User"
	postgres_handler "github.com/aperezgdev/food-order-api/internal/infrastructure/postgres"
)

type UserPostgresRepository struct {
	log             *log.Logger
	postgresHandler postgres_handler.PostgresHandler
}

func NewUserPostgresRepository(
	log *log.Logger,
	postgresHandler postgres_handler.PostgresHandler,
) UserPostgresRepository {
	return UserPostgresRepository{log, postgresHandler}
}

func (ur *UserPostgresRepository) FindById(id value_object.UserId) (entity.User, error) {
	user := new(entity.User)
	result := ur.postgresHandler.DB.First(user, id)

	if result.Error != nil {
		return *user, result.Error
	}

	return *user, nil
}

func (ur *UserPostgresRepository) Save(user entity.User) error {
	err := ur.postgresHandler.Create(user)
	if err != nil {
		ur.log.Panic("UserPostgresRepository.Save", user)
	}

	return err
}
