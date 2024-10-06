package controller

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	application "github.com/aperezgdev/food-order-api/internal/application/order"
	"github.com/aperezgdev/food-order-api/internal/domain/model"
	domain_errors "github.com/aperezgdev/food-order-api/internal/domain/shared/domain_error"
	order_vo "github.com/aperezgdev/food-order-api/internal/domain/value_object/order"
)

type OrderController struct {
	log                *slog.Logger
	orderCreator       *application.OrderCreator
	orderFinderAll     *application.OrderFinderAll
	orderFinderStatus  *application.OrderFinderStatus
	orderStatusUpdater *application.OrderStatusUpdater
}

type statusUpdateBody struct {
	Status string `json:"status"`
}

func NewOrderController(
	log *slog.Logger,
	orderCreator *application.OrderCreator,
	orderFinderAll *application.OrderFinderAll,
	orderFinderStatus *application.OrderFinderStatus,
	orderStatusUpdater *application.OrderStatusUpdater,
) *OrderController {
	return &OrderController{
		log,
		orderCreator,
		orderFinderAll,
		orderFinderStatus,
		orderStatusUpdater,
	}
}

func (oc *OrderController) StatusUpdater(ctx *gin.Context) {
	id := ctx.Param("id")
	statusUpdateBody := statusUpdateBody{}
	err := ctx.ShouldBind(&statusUpdateBody)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Body is not valid")
	}
	oc.log.Info(
		"OrderController.StatusUpdater - Status",
		slog.Any("status", statusUpdateBody.Status),
	)

	result := oc.orderStatusUpdater.Run(
		order_vo.OrderId(id),
		order_vo.OrderStatus(statusUpdateBody.Status),
	)

	result.Error(func(err error) {
		oc.handlerError(err, ctx)
	}).Ok(func(t *model.Order) {
		ctx.JSON(http.StatusOK, t)
	})
}

func (oc *OrderController) FindByStatus(ctx *gin.Context) {
	status := ctx.Param("status")
	oc.log.Info("OrderController.FindByStatus - Status", slog.Any("status", status))

	result := oc.orderFinderStatus.Run(order_vo.OrderStatus(status))

	result.Error(func(err error) {
		oc.handlerError(err, ctx)
	}).Ok(func(t *[]model.Order) {
		ctx.JSON(http.StatusOK, t)
	})
}

func (oc *OrderController) FindAll(ctx *gin.Context) {
	oc.log.Info("OrderController.FindAll")

	result := oc.orderFinderAll.Run()

	result.Error(func(err error) {
		oc.handlerError(err, ctx)
	}).Ok(func(t *[]model.Order) {
		ctx.JSON(http.StatusOK, t)
	})
}

func (oc *OrderController) Create(ctx *gin.Context) {
	var err error
	order := model.Order{}

	err = ctx.ShouldBind(&order)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Body is not valid")
	}

	oc.log.Info("OrderController.Create - Body value", slog.Any("order", order))

	result := oc.orderCreator.Run(order)

	result.Error(func(err error) {
		oc.handlerError(err, ctx)
	}).Ok(func(t *model.Order) {
		ctx.Status(http.StatusCreated)
	})
}

func (oc *OrderController) handlerError(err error, ctx *gin.Context) {
	oc.log.Error("OrderController - Error has ocurred", slog.Any("error", err))

	if errors.Is(err, domain_errors.NotFound) {
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	ctx.AbortWithStatus(http.StatusInternalServerError)
}
