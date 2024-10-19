package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aperezgdev/food-order-api/internal/infrastructure/http/controller"
)

type OrderPathStatusRouteHandler struct {
	orderController *controller.OrderController
}

func NewOrderPathStatusRouteHandler(
	orderController *controller.OrderController,
) *OrderPathStatusRouteHandler {
	return &OrderPathStatusRouteHandler{orderController}
}

func (*OrderPathStatusRouteHandler) Method() string {
	return http.MethodPatch
}

func (*OrderPathStatusRouteHandler) Pattern() string {
	return "/orders/:id"
}

func (oc *OrderPathStatusRouteHandler) Handler(ctx *gin.Context) {
	oc.orderController.StatusUpdater(ctx)
}
