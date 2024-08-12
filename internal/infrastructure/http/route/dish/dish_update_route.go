package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aperezgdev/food-order-api/internal/infrastructure/http/controller"
)

type DishPutRouteHandler struct {
	dishController *controller.DishController
}

func NewDishPutRouteHandler(dishController *controller.DishController) *DishPutRouteHandler {
	return &DishPutRouteHandler{dishController}
}

func (*DishPutRouteHandler) Method() string {
	return http.MethodPut
}

func (*DishPutRouteHandler) Pattern() string {
	return "/dish/:id"
}

func (dp *DishPutRouteHandler) Handler(ctx *gin.Context) {
	dp.dishController.Update(ctx)
}
