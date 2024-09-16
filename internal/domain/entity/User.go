package entity

import (
	. "github.com/aperezgdev/food-order-api/internal/domain/value_object/User"
)

type User struct {
	UserId UserId    `gorm:"type:uuid;default:gen_random_uuid()"`
	Name   UserName  `db:"name"`
	Email  UserEmail `db:"email"`
}

func NewUser(name UserName, email UserEmail) User {
	return User{NewUserId(), name, email}
}
