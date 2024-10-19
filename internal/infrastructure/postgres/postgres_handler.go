package postgres_handler

import (
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/aperezgdev/food-order-api/env"
)

type PostgresHandler struct {
	DB *gorm.DB
}

func NewPostgresHandler(env env.EnvApp, log *slog.Logger) PostgresHandler {
	conn := "host=" + env.DB_HOST + " user=" + env.DB_USER + " password=" + env.DB_PASSWORD + " dbname=" + env.DB_NAME + " port=" + env.PORT_DB + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return PostgresHandler{db}
}
