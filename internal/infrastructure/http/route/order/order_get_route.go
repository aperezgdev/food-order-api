package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aperezgdev/food-order-api/internal/infrastructure/http/controller"
)

type OrderGetRouteHandler struct {
	orderController *controller.OrderController
}

func NewOrderGetRouteHandler(orderController *controller.OrderController) *OrderGetRouteHandler {
	return &OrderGetRouteHandler{orderController}
}

func (*OrderGetRouteHandler) Method() string {
	return http.MethodGet
}

func (*OrderGetRouteHandler) Pattern() string {
	return "/orders"
}

func (og *OrderGetRouteHandler) Handler(ctx *gin.Context) {
	og.orderController.FindAll(ctx)
}
