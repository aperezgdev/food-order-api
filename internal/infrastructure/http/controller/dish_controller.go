package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	application "github.com/aperezgdev/food-order-api/internal/application/Dish"
	"github.com/aperezgdev/food-order-api/internal/domain/entity"
)

type DishController struct {
	slog        *slog.Logger
	dishCreator *application.DishCreator
}

func NewDishController(slog *slog.Logger, dishCreator *application.DishCreator) *DishController {
	return &DishController{slog, dishCreator}
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
