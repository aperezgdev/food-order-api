package route

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aperezgdev/food-order-api/internal/infrastructure/http/controller"
)

type UserPostRouteHandler struct {
	userController *controller.UserController
}

func NewUserPostRouteHandler(userController *controller.UserController) *UserPostRouteHandler {
	return &UserPostRouteHandler{userController}
}

func (*UserPostRouteHandler) Method() string {
	return http.MethodPost
}

func (*UserPostRouteHandler) Pattern() string {
	return "/user"
}

func (uh *UserPostRouteHandler) Handler(ctx *gin.Context) {
	uh.userController.Create(ctx)
}
