package application

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	vo "github.com/aperezgdev/food-order-api/internal/domain/value_object"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/Order"
	"github.com/aperezgdev/food-order-api/internal/infrastructure/repository"
)

func newTestOrderCreator() *OrderCreator {
	return NewOrderCreator(repository.NewOrderInMemoryRepository(), slog.Default())
}

// Should create a valid order
func TestOrderCreatorNotError(t *testing.T) {
	orderCreator := newTestOrderCreator()

	order := entity.Order{
		Id:        value_object.OrderId("2"),
		Status:    value_object.NEW,
		Dishes:    make([]*entity.Dish, 0),
		CreatedOn: vo.NewCreatedOn(),
	}

	var erro error
	result := orderCreator.Run(order)
	result.Error(func(err error) {
		erro = err
	})

	if erro != nil {
		t.Errorf("TestOrderCreator - OrderCreator is returning an error")
	}
}

// Should create and save a valid order
func TestOrderCreator(t *testing.T) {
	orderRepository := repository.NewOrderInMemoryRepository()
	orderCreator := NewOrderCreator(orderRepository, &slog.Logger{})
	orderFinderAll := NewOrderFinderAll(orderRepository, &slog.Logger{})

	var nOrderBefore int
	resultFinderBefore := orderFinderAll.Run()
	resultFinderBefore.Ok(func(t *[]entity.Order) {
		nOrderBefore = len(*t)
	})

	order := entity.Order{
		Id:        value_object.OrderId("3"),
		Status:    value_object.WORKING_ON,
		Dishes:    make([]*entity.Dish, 0),
		CreatedOn: vo.NewCreatedOn(),
	}

	orderCreator.Run(order)

	var nOrdersAfter int
	resultFinderAfter := orderFinderAll.Run()
	resultFinderAfter.Ok(func(t *[]entity.Order) {
		nOrdersAfter = len(*t)
	})

	fmt.Println(nOrdersAfter)
	fmt.Println(nOrderBefore)

	if nOrdersAfter != nOrderBefore+1 {
		t.Errorf("TestOrderCreator - OrderCreator is not saving order")
	}
}
