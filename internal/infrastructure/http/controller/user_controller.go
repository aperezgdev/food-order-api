package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	application "github.com/aperezgdev/food-order-api/internal/application/User"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
)

type UserController struct {
	log         *slog.Logger
	userCreator *application.UserCreator
}

func NewUserController(log *slog.Logger, userCreator *application.UserCreator) *UserController {
	return &UserController{log, userCreator}
}

func (uc *UserController) Create(ctx *gin.Context) {
	var err error
	user := entity.User{}

	err = ctx.ShouldBind(&user)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Body is not valid")
	}

	uc.log.Info("UserController.Create - Body value", user)

	err = uc.userCreator.Run(&user)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
	}
}
