package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	application "github.com/aperezgdev/food-order-api/internal/application/Dish"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
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
	dishes, err := dc.dishFinderAll.Run()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	ctx.JSON(http.StatusOK, dishes)
}

func (dc *DishController) Create(ctx *gin.Context) {
	var err error
	dish := entity.Dish{}

	err = ctx.ShouldBind(&dish)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Body is not valid")
		return
	}

	err = dc.dishCreator.Run(dish)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
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

	err = dc.dishUpdater.Run(dish)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Body is not valid")
		return
	}
}

func (dc *DishController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := dc.dishRemover.Run(value_object.DishId(id))
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	ctx.Status(http.StatusAccepted)
}
