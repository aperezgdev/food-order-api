package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aperezgdev/food-order-api/internal/infrastructure/http/controller"
)

type UserGetRouteHandler struct {
	userController *controller.UserController
}

func NewUserGetRouteHandler(userController *controller.UserController) *UserGetRouteHandler {
	return &UserGetRouteHandler{userController}
}

func (*UserGetRouteHandler) Method() string {
	return http.MethodGet
}

func (*UserGetRouteHandler) Pattern() string {
	return "/user/:id"
}

func (uh *UserGetRouteHandler) Handler(ctx *gin.Context) {
	uh.userController.Find(ctx)
}
