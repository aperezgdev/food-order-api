package postgres_handler

import (
	"github.com/aperezgdev/food-order-api/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresHandler struct {
	DB *gorm.DB
}

func NewPostgresHandler(env env.EnvApp) PostgresHandler {
	conn := "host=db user=" + env.DB_USER + "password=" + env.DB_PASSWORD + "dbname=" +  env.DB_NAME + " port=" + env.PORT_DB + "sslmode=disable"
	db, err := gorm.Open(postgres.Open(conn))
	if err != nil {
		panic(err)
	}

	return PostgresHandler{db}
}

func (ph *PostgresHandler) Create(obj interface{}) error {
	result := ph.DB.Create(&obj)

	return result.Error
}

func (ph *PostgresHandler) FindAll(obj interface{}) error {
	result := ph.DB.Find(obj)

	return result.Error
}

func (ph *PostgresHandler) DeleteById(obj interface{}, id string) error {
	result := ph.DB.Delete(obj, id)

	return result.Error
}
