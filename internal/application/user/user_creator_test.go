package application

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	vo "github.com/aperezgdev/food-order-api/internal/domain/shared/value_object"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/user"
)

// Should create a user without error
func TestUserCreator(t *testing.T) {
	userRepository := repository.NewMockUserRepository()
	userRepository.On("Save", mock.Anything).Return(nil)
	userCreator := NewUserCreator(userRepository, slog.Default())

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
