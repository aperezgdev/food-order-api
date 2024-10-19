package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aperezgdev/food-order-api/internal/infrastructure/http/controller"
)

type DishDeleteRouteHandler struct {
	dishController *controller.DishController
}

func NewDishDeleteRouteHandler(dishController *controller.DishController) *DishDeleteRouteHandler {
	return &DishDeleteRouteHandler{dishController}
}

func (*DishDeleteRouteHandler) Method() string {
	return http.MethodDelete
}

func (*DishDeleteRouteHandler) Pattern() string {
	return "/dish/:id"
}

func (dd *DishDeleteRouteHandler) Handler(ctx *gin.Context) {
	dd.dishController.Delete(ctx)
}
