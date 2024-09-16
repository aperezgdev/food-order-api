package repository

import (
	"log/slog"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/Dish"
	postgres_handler "github.com/aperezgdev/food-order-api/internal/infrastructure/postgres"
)

type DishPostgresRepository struct {
	slog                *slog.Logger
	gormPostgresHandler postgres_handler.GormPostgresHandler
}

func NewDishPostgresRepository(
	slog *slog.Logger,
	gormPostgresHandler postgres_handler.GormPostgresHandler,
) repository.DishRepository {
	return &DishPostgresRepository{slog, gormPostgresHandler}
}

func (dr *DishPostgresRepository) FindAll() ([]entity.Dish, error) {
	dishes := []entity.Dish{}
	ctx := dr.gormPostgresHandler.DB.Find(&dishes)
	if ctx.Error != nil {
		dr.slog.Error("DishPostgresRepository.FindAll - ", ctx.Error.Error())
		return dishes, ctx.Error
	}

	return dishes, nil
}

func (dr *DishPostgresRepository) Save(dish entity.Dish) error {
	ctx := dr.gormPostgresHandler.DB.Create(&dish)
	if ctx.Error != nil {
		dr.slog.Error("DishPostgresRepository.Save", ctx.Error.Error())
		return ctx.Error
	}

	return nil
}

func (dr *DishPostgresRepository) Update(dish entity.Dish) error {
	ctx := dr.gormPostgresHandler.DB.Create(&dish)
	if ctx.Error != nil {
		dr.slog.Error("DishPostgresRepository.Update", ctx.Error.Error())
		return ctx.Error
	}

	return nil
}

func (dr *DishPostgresRepository) Delete(id value_object.DishId) error {
	ctx := dr.gormPostgresHandler.DB.Delete(&entity.Dish{Id: id})
	if ctx.Error != nil {
		dr.slog.Error("DishPostgresRepository.Delete", ctx.Error.Error())
		return ctx.Error
	}

	return nil
}
