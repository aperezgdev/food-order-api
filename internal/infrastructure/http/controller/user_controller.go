package controller

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aperezgdev/food-order-api/internal/application/user"
	"github.com/aperezgdev/food-order-api/internal/domain/model"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/user"
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
	user := model.User{}

	err = ctx.ShouldBind(&user)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Body is not valid")
	}

	uc.log.Info("UserController.Create - Body value", slog.Any("user", user))

	result := uc.userCreator.Run(&user)

	result.Error(func(err error) {
		uc.handlerError(err, ctx)
	}).Ok(func(t *model.User) {
		ctx.Status(http.StatusCreated)
	})
}

func (uc *UserController) Find(ctx *gin.Context) {
	rawId := ctx.Param("id")

	result := uc.userFinder.Run(value_object.UserId(rawId))

	result.Error(func(err error) {
		uc.handlerError(err, ctx)
	}).Ok(func(t *model.User) {
		ctx.JSON(http.StatusOK, t)
	})
}

func (uc *UserController) handlerError(err error, ctx *gin.Context) {
	uc.log.Error("UserController - Error has ocurred", slog.Any("error", err))

	if errors.Is(err, domain_errors.NotFound) {
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	ctx.AbortWithStatus(http.StatusInternalServerError)
}
