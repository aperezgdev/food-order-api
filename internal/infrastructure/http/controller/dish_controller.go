package controller

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	application "github.com/aperezgdev/food-order-api/internal/application/Dish"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/error"
	value_object "github.com/aperezgdev/food-order-api/internal/domain/value_object/Dish"
)

type DishController struct {
	slog          *slog.Logger
	dishCreator   *application.DishCreator
	dishFinderAll *application.DishFinderAll
	dishRemover   *application.DishRemover
	dishUpdater   *application.DishUpdater
}

func NewDishController(
	slog *slog.Logger,
	dishCreator *application.DishCreator,
	dishFinderAll *application.DishFinderAll,
	dishRemover *application.DishRemover,
	dishUpdater *application.DishUpdater,
) *DishController {
	return &DishController{slog, dishCreator, dishFinderAll, dishRemover, dishUpdater}
}

func (dc *DishController) GetAll(ctx *gin.Context) {
	result := dc.dishFinderAll.Run()

	result.Error(func(err error) {
		dc.handlerError(err, ctx)
	}).Ok(func(t *[]entity.Dish) {
		ctx.JSON(http.StatusOK, t)
	})
}

func (dc *DishController) Create(ctx *gin.Context) {
	var err error
	dish := entity.Dish{}

	err = ctx.ShouldBind(&dish)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Body is not valid")
		return
	}

	result := dc.dishCreator.Run(dish)

	result.Error(func(err error) {
		dc.handlerError(err, ctx)
	}).Ok(func(t *entity.Dish) {
		ctx.Status(http.StatusCreated)
	})
}

func (dc *DishController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var err error
	dish := entity.Dish{Id: value_object.DishId(id)}

	err = ctx.ShouldBind(&dish)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Body is not valid")
		return
	}

	result := dc.dishUpdater.Run(dish)

	result.Error(func(err error) {
		dc.handlerError(err, ctx)
	}).Ok(func(t *entity.Dish) {
		ctx.Status(http.StatusOK)
	})
}

func (dc *DishController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	result := dc.dishRemover.Run(value_object.DishId(id))

	result.Error(func(err error) {
		dc.handlerError(err, ctx)
	}).Ok(func(t *entity.Dish) {
		ctx.Status(http.StatusAccepted)
	})
}

func (dc *DishController) handlerError(err error, ctx *gin.Context) {
	dc.slog.Error("UserController - Error has ocurred", slog.Any("error", err))

	if errors.Is(err, domain_errors.NotFound) {
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	ctx.AbortWithStatus(http.StatusInternalServerError)
}
