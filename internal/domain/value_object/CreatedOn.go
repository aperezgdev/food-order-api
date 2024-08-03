package value_object

import "time"

type CreatedOn time.Time

func NewCreatedOn() CreatedOn {
	return CreatedOn(time.Now())
}
