package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	application "github.com/aperezgdev/food-order-api/internal/application/Dish"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
)

type DishController struct {
	slog          *slog.Logger
	dishCreator   *application.DishCreator
	dishFinderAll *application.DishFinderAll
}

func NewDishController(
	slog *slog.Logger,
	dishCreator *application.DishCreator,
	dishFinderAll *application.DishFinderAll,
) *DishController {
	return &DishController{slog, dishCreator, dishFinderAll}
}

func (dc *DishController) GetAll(ctx *gin.Context) {
	dishes, err := dc.dishFinderAll.Run()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
	}

	ctx.JSON(http.StatusOK, dishes)
}

func (dc *DishController) Create(ctx *gin.Context) {
	var err error
	dish := entity.Dish{}

	err = ctx.ShouldBind(&dish)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Body is not valid")
	}

	err = dc.dishCreator.Run(dish)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
	}
}
