package application

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	"github.com/aperezgdev/food-order-api/internal/domain/repository"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/user"
)

// Should find a user
func TestUserFinder(t *testing.T) {
	userRepository := repository.NewMockUserRepository()
	userRepository.On("FindById", mock.AnythingOfType("user_vo.UserId")).
		Return(model.User{Id: value_object.UserId("1")}, nil)
	userFinder := NewUserFinder(userRepository, slog.Default())
	result := userFinder.Run(value_object.UserId("1"))

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if testError != nil {
		t.Errorf("TestUserFinder - User was not found")
	}

	var userFounded model.User
	result.Ok(func(t *model.User) {
		userFounded = *t
	})

	if userFounded.Id != value_object.UserId("1") {
		t.Errorf("TestUserFinder - User founded is not asked user")
	}
}

// Should not find a user and return an not found error
func TestUserFinderError(t *testing.T) {
	userRepository := repository.NewMockUserRepository()
	userRepository.On("FindById", mock.AnythingOfType("user_vo.UserId")).
		Return(model.User{}, domain_errors.NotFound)
	userFinder := NewUserFinder(userRepository, slog.Default())
	result := userFinder.Run(value_object.UserId("1"))

	var testError error
	result.Error(func(err error) {
		testError = err
	})

	if testError == nil {
		t.Errorf("TestUserFinder - User was found")
	}

	if !errors.Is(testError, domain_errors.NotFound) {
		t.Errorf("TestUserFinderError - Error returned is not an domain error NOT FOUND")
	}
}
