package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aperezgdev/food-order-api/internal/infrastructure/http/controller"
)

type DishPostRouteHandler struct {
	dishController *controller.DishController
}

func NewDishPostRouteHandler(dishController *controller.DishController) *DishPostRouteHandler {
	return &DishPostRouteHandler{dishController}
}

func (*DishPostRouteHandler) Method() string {
	return http.MethodPost
}

func (*DishPostRouteHandler) Pattern() string {
	return "/dish"
}

func (dp *DishPostRouteHandler) Handler(ctx *gin.Context) {
	dp.dishController.Create(ctx)
}
