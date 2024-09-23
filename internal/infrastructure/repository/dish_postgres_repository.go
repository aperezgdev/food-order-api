package repository

import (
	"log/slog"

	"gorm.io/gorm"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/dish"
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

func (dr *DishPostgresRepository) FindAll() ([]model.Dish, error) {
	dishes := []model.Dish{}
	ctx := dr.gormPostgresHandler.DB.Find(&dishes)
	if ctx.Error != nil {
		dr.slog.Error("DishPostgresRepository.FindAll - ", ctx.Error.Error())
		return dishes, domain_errors.Database
	}

	return dishes, nil
}

func (dr *DishPostgresRepository) Find(id value_object.DishId) (model.Dish, error) {
	dish := model.Dish{Id: id}
	ctx := dr.gormPostgresHandler.DB.Take(&dish)

	if ctx.Error == gorm.ErrRecordNotFound {
		return dish, domain_errors.NotFound
	}

	if ctx.Error != nil {
		return dish, domain_errors.Database
	}

	return dish, nil
}

func (dr *DishPostgresRepository) Save(dish model.Dish) error {
	ctx := dr.gormPostgresHandler.DB.Create(&dish)
	if ctx.Error != nil {
		dr.slog.Error("DishPostgresRepository.Save", ctx.Error.Error())
		return domain_errors.Database
	}

	return nil
}

func (dr *DishPostgresRepository) Update(dish model.Dish) error {
	ctx := dr.gormPostgresHandler.DB.Create(&dish)

	if ctx.Error != nil {
		dr.slog.Error("DishPostgresRepository.Update", ctx.Error.Error())
		return domain_errors.Database
	}

	return nil
}

func (dr *DishPostgresRepository) Delete(id value_object.DishId) error {
	ctx := dr.gormPostgresHandler.DB.Delete(&model.Dish{Id: id})
	if ctx.Error != nil {
		dr.slog.Error("DishPostgresRepository.Delete", ctx.Error.Error())
		return ctx.Error
	}

	return nil
}
