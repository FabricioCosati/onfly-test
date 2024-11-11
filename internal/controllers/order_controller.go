package controllers

import (
	"net/http"

	"github.com/FabricioCosati/onfly-test/internal/domain/dto"
	"github.com/FabricioCosati/onfly-test/internal/services"
	cval "github.com/FabricioCosati/onfly-test/internal/validator"
	"github.com/gin-gonic/gin"
)

type OrderController interface {
	CreateOrderService(ctx *gin.Context)
	UpdateOrderStatus(ctx *gin.Context)
	GetOrderById(ctx *gin.Context)
	GetOrders(ctx *gin.Context)
}

type OrderControllerImpl struct {
	Service services.OrderService
}

func (impl *OrderControllerImpl) CreateOrderService(ctx *gin.Context) {
	var order dto.OrderRequestPost

	if err := ctx.ShouldBindBodyWithJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := cval.ValidateRequest(order); err.Errors != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Errors})
		return
	}

	response, err := impl.Service.CreateOrderService(ctx, order)
	if err.Err != nil {
		ctx.JSON(err.Code, gin.H{"error": err.ErrorMessage()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (impl *OrderControllerImpl) UpdateOrderStatus(ctx *gin.Context) {
	var order dto.OrderRequestPatch

	if err := ctx.ShouldBindBodyWithJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := cval.ValidateRequest(order); err.Errors != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	response, err := impl.Service.UpdateOrderStatus(ctx, order)
	if err.Err != nil {
		ctx.JSON(err.Code, gin.H{"error": err.ErrorMessage()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (impl *OrderControllerImpl) GetOrderById(ctx *gin.Context) {
	response, err := impl.Service.GetOrderById(ctx)
	if err.Err != nil {
		ctx.JSON(err.Code, gin.H{"error": err.ErrorMessage()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (impl *OrderControllerImpl) GetOrders(ctx *gin.Context) {
	status := ctx.Query("status")
	response, err := impl.Service.GetOrders(ctx, status)
	if err.Err != nil {
		ctx.JSON(err.Code, gin.H{"error": err.ErrorMessage()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func OrderControllerInit(service services.OrderService) *OrderControllerImpl {
	return &OrderControllerImpl{
		Service: service,
	}
}
