package entity

import (
	"github.com/aperezgdev/food-order-api/internal/domain/value_object"
	. "github.com/aperezgdev/food-order-api/internal/domain/value_object/User"
)

type User struct {
	Id        UserId `gorm:"type:uuid;default:gen_random_uuid()"`
	Name      UserName
	Email     UserEmail
	CreatedOn value_object.CreatedOn `gorm:"default:current_timestamp"`
}

func NewUser(name UserName, email UserEmail) User {
	return User{NewUserId(), name, email, value_object.NewCreatedOn()}
}
