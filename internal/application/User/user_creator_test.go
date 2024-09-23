package application

import (
	"log/slog"
	"testing"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	vo "github.com/aperezgdev/food-order-api/internal/domain/shared/value_object"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/user"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/repository"
)

func newTestUserCreator() *UserCreator {
	return NewUserCreator(repository.NewUserInMemoryRepository(), slog.Default())
}

// Should create a user without error
func TestUserCreator(t *testing.T) {
	userCreator := newTestUserCreator()

	user := model.User{
		Id:        value_object.UserId("2"),
		Name:      value_object.UserName("John"),
		Email:     value_object.UserEmail("john@john.com"),
		CreatedOn: vo.NewCreatedOn(),
	}

	result := userCreator.Run(&user)

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if testError != nil {
		t.Errorf("TestUserCreator - Error has ocurred trying to create a user")
	}
}

// Should create and save a user
func TestUserCreatorAndSave(t *testing.T) {
	userRepository := repository.NewUserInMemoryRepository()
	userCreator := NewUserCreator(userRepository, slog.Default())
	userFinder := NewUserFinder(userRepository, slog.Default())

	user := model.User{
		Id:        value_object.UserId("2"),
		Name:      value_object.UserName("John"),
		Email:     value_object.UserEmail("john@john.com"),
		CreatedOn: vo.NewCreatedOn(),
	}

	userCreator.Run(&user)

	var testError error
	result := userFinder.Run(value_object.UserId("2"))
	result.Error(func(err error) {
		testError = err
	})

	if testError != nil {
		t.Errorf("TestUserCreatorAndSave - User was not found")
	}
}
