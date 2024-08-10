package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	application "github.com/aperezgdev/food-order-api/internal/application/User"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/User"
)

type UserController struct {
	log         *slog.Logger
	userCreator *application.UserCreator
	userFinder  *application.UserFinder
}

func NewUserController(
	log *slog.Logger,
	userCreator *application.UserCreator,
	userFinder *application.UserFinder,
) *UserController {
	return &UserController{log, userCreator, userFinder}
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

func (uc *UserController) Find(ctx *gin.Context) {
	rawId := ctx.Param("id")

	user, err := uc.userFinder.Run(value_object.UserId(rawId))
	if err != nil {
		uc.log.Error("UserController.Find", err)
	}

	if user.UserId == "" {
		ctx.String(http.StatusNotFound, "User not found")
	}

	ctx.JSON(http.StatusOK, user)
}
