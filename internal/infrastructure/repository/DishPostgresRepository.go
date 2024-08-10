package repository

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/Dish"
	postgres_handler "github.com/aperezgdev/food-order-api/internal/infrastructure/postgres"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/postgres/queries"
)

type DishPostgresRepository struct {
	slog            *slog.Logger
	postgresHandler postgres_handler.PostgresHandler
}

func NewDishPostgresRepository(
	slog *slog.Logger,
	postgresHandler postgres_handler.PostgresHandler,
) repository.DishRepository {
	return &DishPostgresRepository{slog, postgresHandler}
}

func (dr *DishPostgresRepository) FindAll() ([]entity.Dish, error) {
	dishes := []entity.Dish{}
	err := dr.postgresHandler.DB.Select(&dishes, queries.DishGetAll)
	if err != nil {
		dr.slog.Error("DishPostgresRepository.FindAll - ", err.Error())
		return dishes, err
	}

	return dishes, nil
}

func (dr *DishPostgresRepository) Save(dish entity.Dish) error {
	_, err := dr.postgresHandler.DB.NamedExec(queries.DishCreate, &dish)
	if err != nil {
		dr.slog.Error("DishPostgresRepository.Save", err.Error())
		return err
	}

	return nil
}

func (dr *DishPostgresRepository) Update(dish entity.Dish) error {
	_, err := dr.postgresHandler.DB.NamedExec(queries.DishUpdate, &dish)
	if err != nil {
		dr.slog.Error("DishPostgresRepository.Update", err.Error())
		return err
	}

	return nil
}

func (dr *DishPostgresRepository) Delete(id value_object.DishId) error {
	_, err := dr.postgresHandler.DB.Exec(queries.DishDelete, string(id))
	if err != nil {
		dr.slog.Error("DishPostgresRepository.Delete", err.Error())
		return err
	}

	return nil
}
