package postgres_handler

import (
	"database/sql"
	"log/slog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/aperezgdev/food-order-api/env"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/postgres/queries"
)

type PostgresHandler struct {
	DB *sqlx.DB
}

func NewPostgresHandler(env env.EnvApp, log *slog.Logger) PostgresHandler {
	conn := "host=" + env.DB_HOST + " user=" + env.DB_USER + " password=" + env.DB_PASSWORD + " dbname=" + env.DB_NAME + " port=" + env.PORT_DB + " sslmode=disable"
	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		panic(err)
	}

	db.MustExec(queries.Uuid_extension)
	db.MustExec(queries.UserSchema)
	db.MustExec(queries.DishSchema)

	return PostgresHandler{db}
}

func (ph *PostgresHandler) Create(insert string, obj interface{}) (sql.Result, error) {
	return ph.DB.Exec(insert, &obj)
}
