package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aperezgdev/food-order-api/internal/infrastructure/http/controller"
)

type OrderGetStatusRouteHandler struct {
	orderController *controller.OrderController
}

func NewOrderGetStatusHandler(
	orderController *controller.OrderController,
) *OrderGetStatusRouteHandler {
	return &OrderGetStatusRouteHandler{orderController}
}

func (*OrderGetStatusRouteHandler) Method() string {
	return http.MethodGet
}

func (*OrderGetStatusRouteHandler) Pattern() string {
	return "/orders/status/:status"
}

func (oc *OrderGetStatusRouteHandler) Handler(ctx *gin.Context) {
	oc.orderController.FindByStatus(ctx)
}
