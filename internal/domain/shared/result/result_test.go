package result

import (
	"testing"

	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
)

func TestErrorResult(t *testing.T) {
	result := ErrorResult[string](domain_errors.NotFound)

	var v string
	result.Ok(func(t *string) {
		v = *t
	})

	if v != "" {
		t.Errorf("TestErrorResult - String value should be default init")
	}

	var testErr error
	result.Error(func(err error) {
		testErr = err
	})

	if testErr == nil {
		t.Errorf("TestErrorResult - Error must be not nil")
	}
}

func TestOkResult(t *testing.T) {
	value := "as"
	result := OkResult(&value)

	var testErr error
	result.Error(func(err error) {
		testErr = *&err
	})

	if testErr != nil {
		t.Errorf("TestOkResult - Error must be null")
	}

	var v string
	result.Ok(func(t *string) {
		v = *t
	})

	if v != value {
		t.Errorf("TestOkResult - Value was altered")
	}
}
