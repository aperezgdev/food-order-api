package postgres_handler

import (
	"fmt"
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/aperezgdev/food-order-api/env"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
)

type GormPostgresHandler struct {
	DB *gorm.DB
}

func NewGormPostgresHandler(env env.EnvApp, log *slog.Logger) GormPostgresHandler {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.DB_HOST,
		env.DB_USER,
		env.DB_PASSWORD,
		env.DB_NAME,
		env.PORT_DB,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("NewGormPostgresHandler - Error trying conect with database")
		panic(err)
	}

	errMigrate := db.AutoMigrate(&entity.Order{}, &entity.User{}, &entity.Dish{})

	if errMigrate != nil {
		log.Error("NewGormPostgresHandler - Error on automigration", slog.Any("error", errMigrate))
	}

	return GormPostgresHandler{db}
}
