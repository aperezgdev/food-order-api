package postgres_handler

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresHandler struct {
	DB *gorm.DB
}

func NewPostgresHandler(conn string) PostgresHandler {
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
