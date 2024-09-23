package application

import (
	"log/slog"
	"testing"

	"github.com/aperezgdev/food-order-api/internal/domain/model"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/user"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/repository"
)

func newTestUserFinder() *UserFinder {
	return NewUserFinder(repository.NewUserInMemoryRepository(), slog.Default())
}

// Should find a user
func TestUserFinder(t *testing.T) {
	userFinder := newTestUserFinder()

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
