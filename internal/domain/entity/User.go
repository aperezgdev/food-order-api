package entity

import (
	. "github.com/aperezgdev/food-order-api/internal/domain/value_object/User"
)

type User struct {
	UserId UserId
	Name   UserName
	Email  UserEmail
}

func NewUser(name UserName, email UserEmail) User {
	return User{NewUserId(), name, email}
}
