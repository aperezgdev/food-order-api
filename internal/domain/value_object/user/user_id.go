package user_vo

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

type UserId string

func NewUserId() UserId {
	return UserId(uuid.New().String())
}

func (id *UserId) Value(value interface{}) (driver.Value, error) {
	return string(*id), nil
}
