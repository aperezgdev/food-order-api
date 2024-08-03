package value_object

import "github.com/google/uuid"

type UserId string

func NewUserId() UserId {
	return UserId(uuid.New().String())
}
