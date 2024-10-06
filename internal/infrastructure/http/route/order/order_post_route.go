package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aperezgdev/food-order-api/internal/infrastructure/http/controller"
)

type OrderPostRouteHandler struct {
	orderController *controller.OrderController
}

func NewOrderPostRouteHandler(orderController *controller.OrderController) *OrderPostRouteHandler {
	return &OrderPostRouteHandler{orderController}
}

func (*OrderPostRouteHandler) Method() string {
	return http.MethodPost
}

func (*OrderPostRouteHandler) Pattern() string {
	return "/orders"
}

func (oc *OrderPostRouteHandler) Handler(ctx *gin.Context) {
	oc.orderController.Create(ctx)
}
