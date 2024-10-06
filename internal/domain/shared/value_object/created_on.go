package value_object

import (
	"database/sql/driver"
	"errors"
	"time"
)

type CreatedOn time.Time

func NewCreatedOn() CreatedOn {
	return CreatedOn(time.Now())
}

func (ct *CreatedOn) Scan(value interface{}) error {
	if value == nil {
		*ct = CreatedOn(time.Time{})
		return nil
	}
	if t, ok := value.(time.Time); ok {
		*ct = CreatedOn(t)
		return nil
	}
	return errors.New("Error trying on value object createdOn")
}

func (ct CreatedOn) Value() (driver.Value, error) {
	return time.Time(ct), nil
}
