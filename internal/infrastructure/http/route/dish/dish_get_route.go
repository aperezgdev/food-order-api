package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aperezgdev/food-order-api/internal/infrastructure/http/controller"
)

type DishGetRouteHandler struct {
	dishController *controller.DishController
}

func NewDishGetRouteHandler(dishController *controller.DishController) *DishGetRouteHandler {
	return &DishGetRouteHandler{dishController}
}

func (*DishGetRouteHandler) Method() string {
	return http.MethodGet
}

func (*DishGetRouteHandler) Pattern() string {
	return "/dish"
}

func (dg *DishGetRouteHandler) Handler(ctx *gin.Context) {
	dg.dishController.GetAll(ctx)
}
